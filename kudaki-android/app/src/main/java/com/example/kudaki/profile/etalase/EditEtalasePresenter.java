package com.example.kudaki.profile.etalase;

import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.retrofit.PostData;
import com.example.kudaki.retrofit.RetrofitClient;

import okhttp3.MultipartBody;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class EditEtalasePresenter implements EditEtalaseContract.Presenter {
    String token;
    EditEtalaseContract.View view;

    public EditEtalasePresenter(EditEtalaseContract.View view, String token) {
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void start() {

    }

    @Override
    public void update(String uuid, String name, String desc, String price, String duration) {
        view.showProgress();
        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
        RequestBody requestBody = new MultipartBody.Builder()
                .setType(MultipartBody.FORM)
                .addFormDataPart("item_uuid", uuid)
                .addFormDataPart("name", name)
                .addFormDataPart("amount", "1")
                .addFormDataPart("color", "black")
                .addFormDataPart("description", desc)
                .addFormDataPart("price", price)
                .addFormDataPart("unit", "pairs")
                .addFormDataPart("unit_of_measurement", "CM")
                .addFormDataPart("photo", "http://google.co.id")
                .addFormDataPart("price_duration", "DAY")
                .addFormDataPart("height", "5")
                .addFormDataPart("length", "5")
                .addFormDataPart("width", "5")
                .build();
        Call<DefaultResponse> call = service.updateStoreItem(token, requestBody);

        call.enqueue(new Callback<DefaultResponse>() {
            @Override
            public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                if (response.code() == 200) {
                    view.showEditSuccess("Berhasil simpan");
                } else {
                    view.showEditFailed("Gagal simpan");
                }
                view.closeProgress();
            }

            @Override
            public void onFailure(Call<DefaultResponse> call, Throwable t) {

            }
        });
    }
}
