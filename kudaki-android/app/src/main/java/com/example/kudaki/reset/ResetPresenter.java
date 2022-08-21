package com.example.kudaki.reset;

import android.content.Context;

import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.model.response.ErrorResponse;
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

public class ResetPresenter implements ResetContract.Presenter {
    ResetContract.View view;
    Context context;

    public ResetPresenter(ResetContract.View view, Context context) {
        this.view = view;
        this.context = context;
        this.view.setPresenter(this);
    }

    @Override
    public void start() {

    }

    @Override
    public void doReset(String token, String newPwd) {
        view.showProgress();

        if (newPwd.length() < 8) {
            view.showResetFailed("Gagal daftar! Password Anda kurang dari 8 karakter");
            view.closeProgress();
        } else if (!newPwd.matches("^(?=.*[0-9])(?=.*[a-zA-Z])[a-zA-Z0-9]+$")) {
            view.showResetFailed("Gagal daftar! Password Anda minimal harus memilik 1 angka dan 1 huruf");
            view.closeProgress();
        } else {
            PostData service = RetrofitClient.getRetrofit().create(PostData.class);
            RequestBody requestBody = new MultipartBody.Builder()
                    .setType(MultipartBody.FORM)
                    .addFormDataPart("new_password", newPwd)
                    .build();
            Call<DefaultResponse> call = service.resetPwd(token, requestBody);

            call.enqueue(new Callback<DefaultResponse>() {
                @Override
                public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                    if (response.code() == 200){
                        view.showResetSuccess("Berhasil reset password");
                    }else {
                        Gson gson = new GsonBuilder().create();
                        ErrorResponse errors;
                        try {
                            errors = gson.fromJson(response.errorBody().string(), ErrorResponse.class);
                            view.showResetFailed("Gagal reset password " + errors.getErrors().get(0));
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
}
