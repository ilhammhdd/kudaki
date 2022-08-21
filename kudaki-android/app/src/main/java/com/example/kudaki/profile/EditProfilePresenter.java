package com.example.kudaki.profile;

import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.model.response.ProfileData;
import com.example.kudaki.model.response.ProfileResponse;
import com.example.kudaki.retrofit.GetData;
import com.example.kudaki.retrofit.PostData;
import com.example.kudaki.retrofit.RetrofitClient;

import okhttp3.MultipartBody;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class EditProfilePresenter implements EditProfileContract.Presenter {
    String token;

    EditProfileContract.View view;

    public EditProfilePresenter(EditProfileContract.View view, String token) {
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void start() {

    }

    @Override
    public void loadProfile() {
        view.showProgress();
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<ProfileResponse> call = service.getProfile(token);

        call.enqueue(new Callback<ProfileResponse>() {
            @Override
            public void onResponse(Call<ProfileResponse> call, Response<ProfileResponse> response) {
                if (response.body() != null) {
                    ProfileResponse resp = response.body();

                    ProfileData data = resp.getData(); // simpan data.getToken di cache
                    view.showProfileData(data);
                    view.closeProgress();
                }
            }

            @Override
            public void onFailure(Call<ProfileResponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void update(String name, String phone) {
        view.showProgress();
        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
        RequestBody requestBody = new MultipartBody.Builder()
                .setType(MultipartBody.FORM)
                .addFormDataPart("photo", "http://google.co.id")
                .addFormDataPart("full_name", name)
                .addFormDataPart("phone_number", phone)
                .build();
        Call<DefaultResponse> call = service.updateProfile(token, requestBody);

        call.enqueue(new Callback<DefaultResponse>() {
            @Override
            public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                if (response.code() == 200) {
                     view.showEditSuccess("Berhasil simpan");
                } else {
                    view.showEditFailed("Gagal simpan");
                }
                view.closeProgress();
            }

            @Override
            public void onFailure(Call<DefaultResponse> call, Throwable t) {

            }
        });
    }
}
