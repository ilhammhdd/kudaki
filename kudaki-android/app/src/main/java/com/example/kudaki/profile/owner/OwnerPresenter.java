package com.example.kudaki.profile.owner;

import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.model.response.OwnerHistoryData;
import com.example.kudaki.model.response.OwnerHistoryReponse;
import com.example.kudaki.retrofit.GetData;
import com.example.kudaki.retrofit.PostData;
import com.example.kudaki.retrofit.RetrofitClient;

import okhttp3.MultipartBody;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class OwnerPresenter implements OwnerContract.Presenter {
    String token;
    OwnerContract.View view;

    public OwnerPresenter(OwnerContract.View view, String token) {
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void start() {

    }

    @Override
    public void addItem(String name, String desc, String price, String duration) {
        view.showProgress();
        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
        RequestBody requestBody = new MultipartBody.Builder()
                .setType(MultipartBody.FORM)
                .addFormDataPart("name", name)
                .addFormDataPart("amount", "1")
                .addFormDataPart("color", "black")
                .addFormDataPart("description", desc)
                .addFormDataPart("price", price)
                .addFormDataPart("unit", "pairs")
                .addFormDataPart("unit_of_measurement", "CM")
                .addFormDataPart("photo", "http://google.co.id")
                .addFormDataPart("price_duration", duration)
                .addFormDataPart("height", "5")
                .addFormDataPart("length", "5")
                .addFormDataPart("width", "5")
                .build();
        Call<DefaultResponse> call = service.addStoreItem(token, requestBody);

        call.enqueue(new Callback<DefaultResponse>() {
            @Override
            public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                if (response.code() == 200) {
                    view.showAddSuccess("Berhasil simpan");
                }
                view.closeProgress();
            }

            @Override
            public void onFailure(Call<DefaultResponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void loadPendingNumber() {
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<OwnerHistoryReponse> call = service.ownerOrderHistory(token, 50, 0, "PENDING");

        call.enqueue(new Callback<OwnerHistoryReponse>() {
            @Override
            public void onResponse(Call<OwnerHistoryReponse> call, Response<OwnerHistoryReponse> response) {
                if (response.code() == 200) {
                    OwnerHistoryReponse resp = response.body();

                    OwnerHistoryData data = resp.getData();
                    if (data.getOrders() != null) {
                        view.showPending(data.getOrders().size());
                    }
                }
            }

            @Override
            public void onFailure(Call<OwnerHistoryReponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void loadApprovedNumber() {
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<OwnerHistoryReponse> call = service.ownerOrderHistory(token, 50, 0, "APPROVED");

        call.enqueue(new Callback<OwnerHistoryReponse>() {
            @Override
            public void onResponse(Call<OwnerHistoryReponse> call, Response<OwnerHistoryReponse> response) {
                if (response.code() == 200) {
                    OwnerHistoryReponse resp = response.body();

                    OwnerHistoryData data = resp.getData();
                    if (data.getOrders() != null) {
                        view.showApproved(data.getOrders().size());
                    }
                }
            }

            @Override
            public void onFailure(Call<OwnerHistoryReponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void loadRentedNumber() {
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<OwnerHistoryReponse> call = service.ownerOrderHistory(token, 50, 0, "RENTED");

        call.enqueue(new Callback<OwnerHistoryReponse>() {
            @Override
            public void onResponse(Call<OwnerHistoryReponse> call, Response<OwnerHistoryReponse> response) {
                if (response.code() == 200) {
                    OwnerHistoryReponse resp = response.body();

                    OwnerHistoryData data = resp.getData();
                    if (data.getOrders() != null) {
                        view.showRented(data.getOrders().size());
                    }
                }
            }

            @Override
            public void onFailure(Call<OwnerHistoryReponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void loadDoneNumber() {
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<OwnerHistoryReponse> call = service.ownerOrderHistory(token, 50, 0, "DONE");

        call.enqueue(new Callback<OwnerHistoryReponse>() {
            @Override
            public void onResponse(Call<OwnerHistoryReponse> call, Response<OwnerHistoryReponse> response) {
                if (response.code() == 200) {
                    OwnerHistoryReponse resp = response.body();

                    OwnerHistoryData data = resp.getData();
                    if (data.getOrders() != null) {
                        view.showDone(data.getOrders().size());
                    }
                }
            }

            @Override
            public void onFailure(Call<OwnerHistoryReponse> call, Throwable t) {

            }
        });
    }
}
