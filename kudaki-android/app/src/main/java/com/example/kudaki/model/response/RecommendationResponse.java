package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class RecommendationResponse {
    @SerializedName("data")
    @Expose
    private RecommendationData data;

    public RecommendationData getData() {
        return data;
    }

    public void setData(RecommendationData data) {
        this.data = data;
    }
}
