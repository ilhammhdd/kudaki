package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

public class RecommendationData {
    @SerializedName("recommended_gear")
    @Expose
    private RecommendedGear recommendedGear;
    @SerializedName("recommended_gear_items")
    @Expose
    private ArrayList<RecommendedGearItem> recommendedGearItems = null;
    @SerializedName("recommended_gears")
    @Expose
    private ArrayList<RecommendedGear> recommendedGears = null;

    public ArrayList<RecommendedGear> getRecommendedGears() {
        return recommendedGears;
    }

    public void setRecommendedGears(ArrayList<RecommendedGear> recommendedGears) {
        this.recommendedGears = recommendedGears;
    }

    public RecommendedGear getRecommendedGear() {
        return recommendedGear;
    }

    public void setRecommendedGear(RecommendedGear recommendedGear) {
        this.recommendedGear = recommendedGear;
    }

    public ArrayList<RecommendedGearItem> getRecommendedGearItems() {
        return recommendedGearItems;
    }

    public void setRecommendedGearItems(ArrayList<RecommendedGearItem> recommendedGearItems) {
        this.recommendedGearItems = recommendedGearItems;
    }
}
