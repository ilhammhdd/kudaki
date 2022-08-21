package com.example.kudaki.explore;

import com.example.kudaki.model.request.RecommendationRequest;
import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.model.response.RecommendationData;
import com.example.kudaki.model.response.RecommendationResponse;
import com.example.kudaki.retrofit.GetData;
import com.example.kudaki.retrofit.PostData;
import com.example.kudaki.retrofit.RetrofitClient;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class RecommendationPresenter implements RecommendationContract.Presenter {
    String token, uuid;
    RecommendationContract.View view;

    public RecommendationPresenter(RecommendationContract.View view, String token, String uuid) {
        this.uuid = uuid;
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void loadRecommendation() {
        view.showProgress();
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<RecommendationResponse> call = service.getAllRecommendation(token, uuid, 10, 0);

        call.enqueue(new Callback<RecommendationResponse>() {
            @Override
            public void onResponse(Call<RecommendationResponse> call, Response<RecommendationResponse> response) {
                if (response.code() == 200) {
                    RecommendationResponse resp = response.body();

                    RecommendationData data = resp.getData();
                    view.showData(data);
                }
                view.closeProgress();
            }

            @Override
            public void onFailure(Call<RecommendationResponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void add(RecommendationRequest request) {
        view.showProgress();
        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
        Call<DefaultResponse> call = service.addRecommendation(token, "application/json", request);

        call.enqueue(new Callback<DefaultResponse>() {
            @Override
            public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                if (response.code() == 200) {
                    view.showAddSuccess("Berhasil Tambah");
                } else {
                    view.showAddFailed("Gagal Tambah");
                }
                view.closeProgress();
            }

            @Override
            public void onFailure(Call<DefaultResponse> call, Throwable t) {

            }
        });
    }


    @Override
    public void start() {

    }
}
