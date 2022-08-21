package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class MountainResponse {
    @SerializedName("data")
    @Expose
    private MountainData data;

    public MountainData getData() {
        return data;
    }

    public void setData(MountainData data) {
        this.data = data;
    }
}
