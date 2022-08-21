package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

public class MountainData {
    @SerializedName("mountains")
    @Expose
    private ArrayList<Mountain> mountains = null;

    public ArrayList<Mountain> getMountains() {
        return mountains;
    }

    public void setMountains(ArrayList<Mountain> mountains) {
        this.mountains = mountains;
    }
}
