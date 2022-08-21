package com.example.kudaki.register;

import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.model.response.ErrorResponse;
import com.example.kudaki.model.user.User;
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

public class RegisterPresenter implements RegisterContract.Presenter {
    private RegisterContract.View view;

    public RegisterPresenter(RegisterContract.View registerView) {
        this.view = registerView;
        this.view.setPresenter(this);
    }

    @Override
    public void doRegister(User user) {
        view.showProgress();

        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
        RequestBody requestBody = new MultipartBody.Builder()
                .setType(MultipartBody.FORM)
                .addFormDataPart("full_name", user.getFullname())
                .addFormDataPart("phone_number", user.getPhone())
                .addFormDataPart("email", user.getEmail())
                .addFormDataPart("password", user.getPassword())
                .addFormDataPart("role", "USER")
                .build();
        Call<DefaultResponse> call = service.registerUser(requestBody);

        // validate password
        if (user.getPassword().length() < 8) {
            view.showOnRegisterFailed("Gagal daftar! Password Anda kurang dari 8 karakter");
            view.closeProgress();
        } else if (!user.getPassword().matches("^(?=.*[0-9])(?=.*[a-zA-Z])[a-zA-Z0-9]+$")) {
            view.showOnRegisterFailed("Gagal daftar! Password Anda minimal harus memilik 1 angka dan 1 huruf");
            view.closeProgress();
        } else {
            // if pass, then create user
            call.enqueue(new Callback<DefaultResponse>() {
                @Override
                public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                    if (response.code() == 200) {
                        view.showOnRegisterSuccess("Berhasil daftar! Silahkan cek email untuk verifikasi.");
                    } else {
                        Gson gson = new GsonBuilder().create();
                        ErrorResponse errors;
                        try {
                            errors = gson.fromJson(response.errorBody().string(), ErrorResponse.class);
                            if (errors.getErrors().get(0).contains("email already exists")) {
                                view.showOnRegisterFailed("Email sudah terdaftar");
                            } else if (errors.getErrors().get(0).contains("verification email")) {
                                view.showOnRegisterFailed("Gagal mengirim email verifikasi");
                            } else {
                                view.showOnRegisterFailed("Gagal daftar!" + errors.getErrors().get(0));
                            }
                        } catch (IOException e) {
                            e.printStackTrace();
                        }
                    }
                    view.closeProgress();
                }

                @Override
                public void onFailure(Call<DefaultResponse> call, Throwable t) {

                }
            });
        }
    }

    @Override
    public void backToLogin() {
        view.showLoginActivity();
    }

    @Override
    public boolean validatePassword(String password, String confirmPassword) {
        return password.equals(confirmPassword);
    }

    @Override
    public void start() {

    }
}
