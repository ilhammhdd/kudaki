package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	"github.com/ilhammhdd/kudaki-user-auth-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-user-auth-service/externals"
	kudakigrpc "github.com/ilhammhdd/kudaki-user-auth-service/externals/grpc"
	"github.com/ilhammhdd/kudaki-user-auth-service/externals/mysql"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

func init() {
	if len(os.Args) > 1 {
		for _, val := range os.Args[1:] {
			f := strings.Split(val, " ")
			os.Setenv(string(f[1]), f[2])
		}
	}

	mysql.CommandDB = mysql.OpenDB(os.Getenv("DB_PATH"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	mysql.QueryDB = mysql.OpenDB(os.Getenv("QUERY_DB_PATH"), os.Getenv("QUERY_DB_USERNAME"), os.Getenv("QUERY_DB_PASSWORD"), os.Getenv("QUERY_DB_NAME"))

	initJWT()
	initAdmin()
}

func initJWT() {
	e := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}
	errorkit.ErrorHandled(jwtkit.GeneratePublicPrivateToPEM(e))

	ecdsa := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("RESET_PASSWORD_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("RESET_PASSWORD_PUBLIC_KEY")}
	errorkit.ErrorHandled(jwtkit.GeneratePublicPrivateToPEM(ecdsa))
}

func initAdmin() {
	if adminExists() {
		log.Println("admin already exists")
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte("OlahragaOtak2K19!"), bcrypt.MinCost)
	errorkit.ErrorHandled(err)

	admin := &user.User{
		AccountType: user.AccountType_NATIVE,
		Email:       "kudaki.service@gmail.com",
		Password:    string(password),
		PhoneNumber: "",
		Role:        user.UserRole_ADMIN,
		Token:       "",
		Uuid:        uuid.New().String()}

	if adminExists() {
		log.Println("admin already exists")
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err = dbo.Command(
		"INSERT INTO kudaki_user.users(uuid,email,password,token,role,phone_number,account_type) VALUES(?,?,?,?,?,?,?)",
		admin.Uuid, admin.Email, admin.Password, admin.Token, admin.Role.String(), admin.PhoneNumber, admin.AccountType.String())
	errorkit.ErrorHandled(err)

	dboProfile := mysql.NewDBOperation(mysql.CommandDB)
	_, err = dboProfile.Command(
		"INSERT INTO kudaki_user.profiles(uuid,user_uuid,full_name,created_at) VALUES(?,?,?,UNIX_TIMESTAMP())",
		uuid.New().String(), admin.Uuid, "Administrator")
	errorkit.ErrorHandled(err)
}

func adminExists() bool {
	dbo := mysql.NewDBOperation(mysql.QueryDB)
	row, err := dbo.QueryRow(
		"SELECT count(id) FROM kudaki_user.users WHERE role=?",
		user.UserRole_ADMIN.String())
	errorkit.ErrorHandled(err)

	var totalIds int
	row.Scan(&totalIds)

	return totalIds == 1
}

func grpcListener() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	errorkit.ErrorHandled(err)

	grpcServer := grpc.NewServer()
	kudakigrpc.RegisterUserServer(grpcServer, externals.UsergRPCService{})
	errorkit.ErrorHandled(grpcServer.Serve(lis))
}

func main() {
	wp := safekit.NewWorkerPool()

	wp.Work <- grpcListener
	wp.Worker <- new(externals.Login)
	wp.Worker <- new(externals.Signup)
	wp.Worker <- new(externals.VerifyUser)
	wp.Worker <- new(externals.ChangePassword)
	wp.Worker <- new(externals.ResetPasswordSendEmail)
	wp.Worker <- new(externals.ResetPassword)

	wp.PoolWG.Wait()
}
