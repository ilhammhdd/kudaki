package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class AllItemResponse {
    @SerializedName("data")
    @Expose
    private AllItemData data;

    public AllItemData getData() {
        return data;
    }

    public void setData(AllItemData data) {
        this.data = data;
    }
}
