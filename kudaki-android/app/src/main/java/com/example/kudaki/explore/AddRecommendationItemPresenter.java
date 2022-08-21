package com.example.kudaki.explore;

import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.retrofit.PostData;
import com.example.kudaki.retrofit.RetrofitClient;

import okhttp3.MultipartBody;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class AddRecommendationItemPresenter implements AddRecommendationItemContract.Presenter {
    String token, uuid;
    AddRecommendationItemContract.View view;

    public AddRecommendationItemPresenter(AddRecommendationItemContract.View view, String token, String uuid) {
        this.uuid = uuid;
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void add(String item, String total) {
        view.showProgress();
        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
        RequestBody requestBody = new MultipartBody.Builder()
                .setType(MultipartBody.FORM)
                .addFormDataPart("recommended_gear_uuid", uuid)
                .addFormDataPart("item_type", item)
                .addFormDataPart("total", total)
                .build();
        Call<DefaultResponse> call = service.addRecommendationItem(token, requestBody);

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
