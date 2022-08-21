package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class StoreResponse {
    @SerializedName("data")
    @Expose
    private StoreData data;

    public StoreData getData() {
        return data;
    }

    public void setData(StoreData data) {
        this.data = data;
    }
}
