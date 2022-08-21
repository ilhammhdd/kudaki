package com.example.kudaki.cart;

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

public class CartPresenter implements CartContract.Presenter {
    String token;
    CartContract.View view;

    public CartPresenter(CartContract.View view, String token) {
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void loadItems() {
        GetData service = RetrofitClient.getRetrofit().create(GetData.class);
        Call<CartResponse> call = service.getCartItems(token, 0, 10);

        call.enqueue(new Callback<CartResponse>() {
            @Override
            public void onResponse(Call<CartResponse> call, Response<CartResponse> response) {
                if (response.code() == 200) {
                    CartResponse resp = response.body();

                    CartData data = resp.getData();
                    view.showCartItems(data);
                }
            }

            @Override
            public void onFailure(Call<CartResponse> call, Throwable t) {

            }
        });
    }

    @Override
    public void checkout(String uuid) {
        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
        RequestBody requestBody = new MultipartBody.Builder()
                .setType(MultipartBody.FORM)
                .addFormDataPart("cart_uuid", uuid)
                .build();
        Call<DefaultResponse> call = service.checkout(token, requestBody);

        call.enqueue(new Callback<DefaultResponse>() {
            @Override
            public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                if (response.code() == 200) {
                    view.showCheckoutSuccess("Berhasil checkout");
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
