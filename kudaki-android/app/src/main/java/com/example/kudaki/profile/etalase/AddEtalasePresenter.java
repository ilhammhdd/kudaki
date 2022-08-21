package com.example.kudaki.profile.etalase;

import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.retrofit.PostData;
import com.example.kudaki.retrofit.RetrofitClient;

import okhttp3.MultipartBody;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class AddEtalasePresenter implements AddEtalaseContract.Presenter {
    String token;
    AddEtalaseContract.View view;

    String filePath;

    public AddEtalasePresenter(AddEtalaseContract.View view, String token) {
        this.token = token;
        this.view = view;
        this.view.setPresenter(this);
    }

    @Override
    public void start() {

    }

    @Override
    public void addItem(String photo, String name, String desc, String price, String amount) {
        view.showProgress();

        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
        RequestBody requestBody = new MultipartBody.Builder()
                .setType(MultipartBody.FORM)
                .addFormDataPart("name", name)
                .addFormDataPart("amount", amount)
                .addFormDataPart("color", "black")
                .addFormDataPart("description", desc)
                .addFormDataPart("price", price)
                .addFormDataPart("unit", "pairs")
                .addFormDataPart("unit_of_measurement", "CM")
                .addFormDataPart("photo", photo)
                .addFormDataPart("price_duration", "DAY")
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

//        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
//        RequestBody uploadBody = new MultipartBody.Builder()
//                .setType(MultipartBody.FORM)
//                .addFormDataPart("file", photo)
//                .build();
//
//        Call<FileResponse> upload = service.uploadFile(uploadBody);
//        upload.enqueue(new Callback<FileResponse>() {
//            @Override
//            public void onResponse(Call<FileResponse> call, Response<FileResponse> response) {
//                if (response.code() == 200) {
//                    FileResponse resp = response.body();
//
//                    FileData data = resp.getData();
//
//                    filePath = data.getFullPath();
//                    if (!filePath.isEmpty()) {
//                        submitItem(name, desc, price, amount);
//                    }
//                } else {
//                    view.showAddFailed("Gagal Unggah Foto");
//                    view.closeProgress();
//                }
//            }
//
//            @Override
//            public void onFailure(Call<FileResponse> call, Throwable t) {
//
//            }
//        });
    }

    private void submitItem(String name, String desc, String price, String amount) {
        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
        RequestBody requestBody = new MultipartBody.Builder()
                .setType(MultipartBody.FORM)
                .addFormDataPart("name", name)
                .addFormDataPart("amount", amount)
                .addFormDataPart("color", "black")
                .addFormDataPart("description", desc)
                .addFormDataPart("price", price)
                .addFormDataPart("unit", "pairs")
                .addFormDataPart("unit_of_measurement", "CM")
                .addFormDataPart("photo", filePath)
                .addFormDataPart("price_duration", "DAY")
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
}
