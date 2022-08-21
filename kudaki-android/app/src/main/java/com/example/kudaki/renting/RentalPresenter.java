package com.example.kudaki.renting;

import com.example.kudaki.model.response.AllItemData;
import com.example.kudaki.model.response.AllItemResponse;
import com.example.kudaki.retrofit.GetData;
import com.example.kudaki.retrofit.RetrofitClient;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class RentalPresenter implements RentalContract.Presenter {
    String token;
    RentalContract.View view;

    public RentalPresenter(RentalContract.View view, String token) {
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void loadItems() {
        view.showProgress();
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<AllItemResponse> call = service.getAllItems(token, 0, 10);

        call.enqueue(new Callback<AllItemResponse>() {
            @Override
            public void onResponse(Call<AllItemResponse> call, Response<AllItemResponse> response) {
                if (response.code() == 200) {
                    AllItemResponse resp = response.body();

                    AllItemData data = resp.getData();
                    view.displayItems(data);
                }
                view.closeProgress();
            }

            @Override
            public void onFailure(Call<AllItemResponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void start() {

    }
}
