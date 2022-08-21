package com.example.kudaki.transaction;

import com.example.kudaki.model.response.OwnerHistoryData;
import com.example.kudaki.model.response.OwnerHistoryReponse;
import com.example.kudaki.retrofit.GetData;
import com.example.kudaki.retrofit.RetrofitClient;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class OwnerTransactionPresenter implements OwnerTransactionContract.Presenter {
    String token;
    OwnerTransactionContract.View view;

    public OwnerTransactionPresenter(OwnerTransactionContract.View view, String token) {
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void loadTransaction(String status) {
        view.showProgress();
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<OwnerHistoryReponse> call = service.ownerOrderHistory(token, 15, 0, status);

        call.enqueue(new Callback<OwnerHistoryReponse>() {
            @Override
            public void onResponse(Call<OwnerHistoryReponse> call, Response<OwnerHistoryReponse> response) {
                if (response.code() == 200) {
                    OwnerHistoryReponse resp = response.body();

                    OwnerHistoryData data = resp.getData();
                    view.showOrderHistoryData(data);
                }
                view.closeProgress();
            }

            @Override
            public void onFailure(Call<OwnerHistoryReponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void start() {

    }
}
