package com.example.kudaki.login;

import android.util.Log;

import com.example.kudaki.model.response.ErrorResponse;
import com.example.kudaki.model.response.LoginData;
import com.example.kudaki.model.response.LoginResponse;
import com.example.kudaki.retrofit.PostData;
import com.example.kudaki.retrofit.RetrofitClient;
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

import java.io.IOException;

import okhttp3.MultipartBody;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class LoginPresenter implements LoginContract.Presenter {
    private LoginContract.View view;

    public LoginPresenter(LoginContract.View loginView) {
        this.view = loginView;
        this.view.setPresenter(this);
    }

    @Override
    public void doLogin(String email, String password) {
        view.showProgress();

        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
        RequestBody requestBody = new MultipartBody.Builder()
                .setType(MultipartBody.FORM)
                .addFormDataPart("email", email)
                .addFormDataPart("password", password)
                .build();
        Call<LoginResponse> call = service.loginUser(requestBody);

        call.enqueue(new Callback<LoginResponse>() {

            @Override
            public void onResponse(Call<LoginResponse> call, Response<LoginResponse> response) {
                if (response.body() != null) {
                    LoginResponse resp = response.body();

                    LoginData data = resp.getData(); // simpan data.getToken di cache
                    Log.d("LoginPresenter", "onResponse: Token = "+ data.getToken());
                    view.showOnLoginSuccess("Berhasil Login!", data.getToken());
                } else if (response.code() >= 500) {
                    view.showOnLoginFailed("Terjadi kesalahan pada server. Silahkan coba kembali.");
                } else {
                    Gson gson = new GsonBuilder().create();
                    ErrorResponse errors;
                    try {
                        errors = gson.fromJson(response.errorBody().string(), ErrorResponse.class);
                        if (errors.getErrors().get(0).contains("email doesn't exists")) {
                            view.showOnLoginFailed("Email tidak terdaftar");
                        } else if (errors.getErrors().get(0).contains("wrong password")) {
                            view.showOnLoginFailed("Password salah");
                        } else if (errors.getErrors().get(0).contains("user wasn't verified")) {
                            view.showOnLoginFailed("Akun belum diverifikasi");
                        } else {
                            view.showOnLoginFailed("Gagal masuk!");
                        }
                    } catch (IOException e) {
                        e.printStackTrace();
                    }
                }
                view.closeProgress();
            }

            @Override
            public void onFailure(Call<LoginResponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void start() {

    }
}
