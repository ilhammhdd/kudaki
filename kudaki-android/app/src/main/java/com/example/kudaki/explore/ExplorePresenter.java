package com.example.kudaki.explore;

import com.example.kudaki.model.response.MountainData;
import com.example.kudaki.model.response.MountainResponse;
import com.example.kudaki.retrofit.GetData;
import com.example.kudaki.retrofit.RetrofitClient;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class ExplorePresenter implements ExploreContract.Presenter {
    ExploreContract.View view;
    String token;

    public ExplorePresenter(ExploreContract.View view, String token) {
        this.view = view;
        this.token = token;
        this.view.setPresenter(this);
    }

    @Override
    public void loadMountain() {
        view.showProgress();
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<MountainResponse> call = service.getAllMountain(token, 15, 0);

        call.enqueue(new Callback<MountainResponse>() {
            @Override
            public void onResponse(Call<MountainResponse> call, Response<MountainResponse> response) {
                if (response.code() == 200) {
                    MountainResponse resp = response.body();

                    MountainData data = resp.getData();
                    view.showMountainData(data);
                }

                view.closeProgress();
            }

            @Override
            public void onFailure(Call<MountainResponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void start() {

    }
}
