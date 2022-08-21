package com.example.kudaki;

import com.example.kudaki.model.response.MountainData;
import com.example.kudaki.model.response.MountainResponse;
import com.example.kudaki.retrofit.GetData;
import com.example.kudaki.retrofit.RetrofitClient;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class MainPresenter implements MainContract.Presenter {
    String token;
    MainContract.View view;

    public MainPresenter(MainContract.View view, String token) {
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void start() {

    }

    @Override
    public void loadPopular() {
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<MountainResponse> call = service.getAllMountain(token, 5, 0);

        call.enqueue(new Callback<MountainResponse>() {
            @Override
            public void onResponse(Call<MountainResponse> call, Response<MountainResponse> response) {
                if (response.code() == 200) {
                    MountainResponse resp = response.body();

                    MountainData data = resp.getData();
                    view.showPopularData(data);
                }
            }

            @Override
            public void onFailure(Call<MountainResponse> call, Throwable t) {

            }
        });
    }
}
