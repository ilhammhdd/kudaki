package com.example.kudaki.profile;

import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.retrofit.PostData;
import com.example.kudaki.retrofit.RetrofitClient;

import okhttp3.MultipartBody;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class EditPasswordPresenter implements EditPasswordContract.Presenter {
    String token;

    EditPasswordContract.View view;

    public EditPasswordPresenter(EditPasswordContract.View view, String token) {
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void update(String oldPwd, String newPwd) {
        view.showProgress();
        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
        RequestBody requestBody = new MultipartBody.Builder()
                .setType(MultipartBody.FORM)
                .addFormDataPart("new_password", newPwd)
                .addFormDataPart("old_password", oldPwd)
                .build();
        Call<DefaultResponse> call = service.changePwd(token, requestBody);

        // validate password
        if (newPwd.length() < 8) {
            view.showEditFailed("Gagal daftar! Password Anda kurang dari 8 karakter");
            view.closeProgress();
        } else if (!newPwd.matches("^(?=.*[0-9])(?=.*[a-zA-Z])[a-zA-Z0-9]+$")) {
            view.showEditFailed("Gagal daftar! Password Anda minimal harus memilik 1 angka dan 1 huruf");
            view.closeProgress();
        } else {
            call.enqueue(new Callback<DefaultResponse>() {
                @Override
                public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                    if (response.code() == 200) {
                        view.showEditSuccess("Berhasil simpan");
                        view.closeProgress();
                    }
                }

                @Override
                public void onFailure(Call<DefaultResponse> call, Throwable t) {

                }
            });
        }
    }

    @Override
    public void start() {

    }
}
