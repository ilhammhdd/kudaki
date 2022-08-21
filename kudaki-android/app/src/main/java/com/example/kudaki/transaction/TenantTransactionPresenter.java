package com.example.kudaki.transaction;

import com.example.kudaki.model.response.OrderHistoryData;
import com.example.kudaki.model.response.OrderHistoryResponse;
import com.example.kudaki.retrofit.GetData;
import com.example.kudaki.retrofit.RetrofitClient;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class TenantTransactionPresenter implements TenantTransactionContract.Presenter {
    String token;
    TenantTransactionContract.View view;

    public TenantTransactionPresenter(TenantTransactionContract.View view, String token) {
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void start() {

    }

    @Override
    public void loadTransaction(String status) {
        view.showProgress();
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<OrderHistoryResponse> call = service.getOrderHistory(token, 15, 0, status);

        call.enqueue(new Callback<OrderHistoryResponse>() {
            @Override
            public void onResponse(Call<OrderHistoryResponse> call, Response<OrderHistoryResponse> response) {
                if (response.code() == 200) {
                    OrderHistoryResponse resp = response.body();

                    OrderHistoryData data = resp.getData();
                    view.showOrderHistoryData(data);
                }
                view.closeProgress();
            }

            @Override
            public void onFailure(Call<OrderHistoryResponse> call, Throwable t) {

            }
        });
    }
}
