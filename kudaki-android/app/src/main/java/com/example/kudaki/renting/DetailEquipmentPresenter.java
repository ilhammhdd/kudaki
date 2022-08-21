package com.example.kudaki.renting;

import com.example.kudaki.model.response.CartData;
import com.example.kudaki.model.response.CartResponse;
import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.retrofit.GetData;
import com.example.kudaki.retrofit.PostData;
import com.example.kudaki.retrofit.RetrofitClient;

import okhttp3.MultipartBody;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class DetailEquipmentPresenter implements DetailEquipmentContract.Presenter {
    String token, uuid, cartUuid;
    DetailEquipmentContract.View view;

    public DetailEquipmentPresenter(DetailEquipmentContract.View view, String token, String uuid) {
        this.token = token;
        this.uuid = uuid;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void addItem() {
        view.showProgress();
        Long time = System.currentTimeMillis() / 1000L;

        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
        RequestBody requestBody = new MultipartBody.Builder()
                .setType(MultipartBody.FORM)
                .addFormDataPart("item_uuid", uuid)
                .addFormDataPart("item_amount", "1")
                .addFormDataPart("duration_from", String.valueOf(time))
                .addFormDataPart("duration", "1")
                .build();

        Call<DefaultResponse> call = service.addToCart(token, requestBody);

        call.enqueue(new Callback<DefaultResponse>() {
            @Override
            public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                if (response.code() == 200) {
                    view.showAddSuccess("Berhasil Ditambahkan");
                }
                view.closeProgress();
            }

            @Override
            public void onFailure(Call<DefaultResponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void checkout() {
        view.showProgress();
        Long time = System.currentTimeMillis() / 1000L;

        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
        RequestBody requestBody = new MultipartBody.Builder()
                .setType(MultipartBody.FORM)
                .addFormDataPart("item_uuid", uuid)
                .addFormDataPart("item_amount", "1")
                .addFormDataPart("duration_from", String.valueOf(time))
                .addFormDataPart("duration", "1")
                .build();

        Call<DefaultResponse> call = service.addToCart(token, requestBody);

        call.enqueue(new Callback<DefaultResponse>() {
            @Override
            public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                if (response.code() == 200) {
                    GetData getService = RetrofitClient.getRetrofit().create(GetData.class);
                    Call<CartResponse> callCart = getService.getCartItems(token, 0, 10);

                    callCart.enqueue(new Callback<CartResponse>() {
                        @Override
                        public void onResponse(Call<CartResponse> call, Response<CartResponse> response) {
                            if (response.code() == 200) {
                                CartResponse resp = response.body();

                                CartData data = resp.getData();
                                cartUuid = data.getCart().getUuid();

                                PostData service = RetrofitClient.getRetrofit().create(PostData.class);
                                RequestBody requestBody = new MultipartBody.Builder()
                                        .setType(MultipartBody.FORM)
                                        .addFormDataPart("cart_uuid", cartUuid)
                                        .build();
                                Call<DefaultResponse> callCheckout = service.checkout(token, requestBody);

                                callCheckout.enqueue(new Callback<DefaultResponse>() {
                                    @Override
                                    public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                                        if (response.code() == 200) {
                                            view.showCheckoutSuccess("Berhasil Checkout");
                                        }
                                        view.closeProgress();
                                    }

                                    @Override
                                    public void onFailure(Call<DefaultResponse> call, Throwable t) {

                                    }
                                });
                            }
                        }

                        @Override
                        public void onFailure(Call<CartResponse> call, Throwable t) {

                        }
                    });
                }
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
