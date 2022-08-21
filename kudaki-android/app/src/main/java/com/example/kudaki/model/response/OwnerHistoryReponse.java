package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class OwnerHistoryReponse {
    @SerializedName("data")
    @Expose
    private OwnerHistoryData data;

    public OwnerHistoryData getData() {
        return data;
    }

    public void setData(OwnerHistoryData data) {
        this.data = data;
    }
}
