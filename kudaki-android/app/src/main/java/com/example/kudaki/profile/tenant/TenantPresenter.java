package com.example.kudaki.profile.tenant;

import com.example.kudaki.model.response.OrderHistoryData;
import com.example.kudaki.model.response.OrderHistoryResponse;
import com.example.kudaki.retrofit.GetData;
import com.example.kudaki.retrofit.RetrofitClient;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class TenantPresenter implements TenantContract.Presenter {
    String token;
    TenantContract.View view;

    public TenantPresenter(TenantContract.View view, String token) {
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void start() {

    }

    @Override
    public void loadPendingNumber() {
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<OrderHistoryResponse> call = service.getOrderHistory(token, 50, 0, "PENDING");

        call.enqueue(new Callback<OrderHistoryResponse>() {
            @Override
            public void onResponse(Call<OrderHistoryResponse> call, Response<OrderHistoryResponse> response) {
                if (response.code() == 200) {
                    OrderHistoryResponse resp = response.body();

                    OrderHistoryData data = resp.getData();
                    if (data.getOrders() != null) {
                        view.showPending(data.getOrders().size());
                    }
                }
            }

            @Override
            public void onFailure(Call<OrderHistoryResponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void loadApprovedNumber() {
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<OrderHistoryResponse> call = service.getOrderHistory(token, 50, 0, "APPROVED");

        call.enqueue(new Callback<OrderHistoryResponse>() {
            @Override
            public void onResponse(Call<OrderHistoryResponse> call, Response<OrderHistoryResponse> response) {
                if (response.code() == 200) {
                    OrderHistoryResponse resp = response.body();

                    OrderHistoryData data = resp.getData();
                    if (data.getOrders() != null) {
                        view.showApproved(data.getOrders().size());
                    }
                }
            }

            @Override
            public void onFailure(Call<OrderHistoryResponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void loadRentedNumber() {
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<OrderHistoryResponse> call = service.getOrderHistory(token, 50, 0, "RENTED");

        call.enqueue(new Callback<OrderHistoryResponse>() {
            @Override
            public void onResponse(Call<OrderHistoryResponse> call, Response<OrderHistoryResponse> response) {
                if (response.code() == 200) {
                    OrderHistoryResponse resp = response.body();

                    OrderHistoryData data = resp.getData();
                    if (data.getOrders() != null) {
                        view.showRented(data.getOrders().size());
                    }
                }
            }

            @Override
            public void onFailure(Call<OrderHistoryResponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void loadDoneNumber() {
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<OrderHistoryResponse> call = service.getOrderHistory(token, 50, 0, "DONE");

        call.enqueue(new Callback<OrderHistoryResponse>() {
            @Override
            public void onResponse(Call<OrderHistoryResponse> call, Response<OrderHistoryResponse> response) {
                if (response.code() == 200) {
                    OrderHistoryResponse resp = response.body();

                    OrderHistoryData data = resp.getData();
                    if (data.getOrders() != null) {
                        view.showDone(data.getOrders().size());
                    }
                }
            }

            @Override
            public void onFailure(Call<OrderHistoryResponse> call, Throwable t) {

            }
        });
    }
}
