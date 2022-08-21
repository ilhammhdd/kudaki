package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class FileResponse {
    @SerializedName("data")
    @Expose
    private FileData data;

    public FileData getData() {
        return data;
    }

    public void setData(FileData data) {
        this.data = data;
    }
}
