package com.example.kudaki.profile;

import com.example.kudaki.model.response.ProfileData;
import com.example.kudaki.model.response.ProfileResponse;
import com.example.kudaki.retrofit.GetData;
import com.example.kudaki.retrofit.RetrofitClient;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class ProfilePresenter implements ProfileContract.Presenter {
    String token;

    ProfileContract.View view;

    public ProfilePresenter(ProfileContract.View view, String token) {
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
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
    public void start() {

    }
}
