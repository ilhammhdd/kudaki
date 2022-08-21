package com.example.kudaki.login;

import android.app.ProgressDialog;
import android.content.Intent;
import android.os.Bundle;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.Nullable;
import androidx.appcompat.app.AppCompatActivity;

import com.example.kudaki.MainActivity;
import com.example.kudaki.R;
import com.example.kudaki.forgotpwd.ForgotPwdActivity;
import com.example.kudaki.register.RegisterActivity;
import com.facebook.login.widget.LoginButton;
import com.orhanobut.hawk.Hawk;

import butterknife.BindView;
import butterknife.ButterKnife;

public class LoginActivity extends AppCompatActivity implements LoginContract.View {

    @BindView(R.id.loginEmail)
    EditText email;
    @BindView(R.id.loginPassword)
    EditText password;
    @BindView(R.id.submitLogin)
    Button button;
    @BindView(R.id.linkSignup)
    TextView linkSignup;
    @BindView(R.id.linkForgotPwd)
    TextView linkForgotPwd;
    @BindView(R.id.buttonLoginGoogle)
    ImageView buttonLoginGoogle;
    @BindView(R.id.buttonLoginFacebook)
    LoginButton loginButton;

    String token;

    ProgressDialog progressDialog;
    LoginPresenter loginPresenter;
    LoginContract.Presenter contractPresenter;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_login);
        ButterKnife.bind(this);

        Hawk.init(this).build();

        token = Hawk.get("token");

        loginPresenter = new LoginPresenter(this);

        progressDialog = new ProgressDialog(this);
    }

    @Override
    protected void onStart() {
        super.onStart();
    }

    @Override
    protected void onResume() {
        super.onResume();
        button.setOnClickListener(v ->
                contractPresenter.doLogin(email.getText().toString(), password.getText().toString()));

        linkSignup.setOnClickListener(v -> {
            Intent signup = new Intent(this, RegisterActivity.class);
            startActivityForResult(signup, 1);
        });

        linkForgotPwd.setOnClickListener(v -> {
            Intent forgot = new Intent(this, ForgotPwdActivity.class);
            startActivity(forgot);
        });
    }

    // on activity result
    @Override
    protected void onActivityResult(int requestCode, int resultCode, @Nullable Intent data) {
        super.onActivityResult(requestCode, resultCode, data);

        // if result code from Register Activity is OK
        if (resultCode == RESULT_OK) {
            this.finish(); // Finish this Login Activity
        }
    }

    @Override
    public void setPresenter(LoginContract.Presenter contractPresenter) {
        this.contractPresenter = contractPresenter;
    }

    @Override
    public void showOnLoginSuccess(String message, String token) {
        closeProgress();
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
        Intent home = new Intent(this, MainActivity.class);

        Hawk.put("token", token);

        startActivity(home);
        finish();
    }

    @Override
    public void showOnLoginFailed(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
    }

    @Override
    public void showProgress() {
        progressDialog.setMax(100);
        progressDialog.setMessage("Please wait...");
        progressDialog.setTitle("Loading");
        progressDialog.setProgressStyle(ProgressDialog.STYLE_SPINNER);
        progressDialog.show();
    }

    @Override
    public void closeProgress() {
        progressDialog.dismiss();
    }
}