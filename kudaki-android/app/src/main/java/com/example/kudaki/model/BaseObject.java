package com.example.kudaki.model;

import com.google.gson.annotations.SerializedName;

public class BaseObject {
    @SerializedName("success")
    boolean success;

    public boolean isSuccess() {
        return success;
    }
}
