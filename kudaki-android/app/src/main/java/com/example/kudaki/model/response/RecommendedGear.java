package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class RecommendedGear {
    @SerializedName("uuid")
    @Expose
    private String uuid;
    @SerializedName("upvote")
    @Expose
    private Integer upvote;
    @SerializedName("downvote")
    @Expose
    private Integer downvote;
    @SerializedName("seen")
    @Expose
    private Integer seen;
    @SerializedName("created_at")
    @Expose
    private Integer createdAt;
    @SerializedName("creator_full_name")
    @Expose
    private String creatorFullName;
    @SerializedName("creator_email")
    @Expose
    private String creatorEmail;
    @SerializedName("mountain_name")
    @Expose
    private String mountainName;
    @SerializedName("mountain_photo")
    @Expose
    private String mountainPhoto;

    public RecommendedGear(String uuid, Integer seen, String creatorFullName) {
        this.uuid = uuid;
        this.seen = seen;
        this.creatorFullName = creatorFullName;
    }

    public String getUuid() {
        return uuid;
    }

    public void setUuid(String uuid) {
        this.uuid = uuid;
    }

    public Integer getUpvote() {
        return upvote;
    }

    public void setUpvote(Integer upvote) {
        this.upvote = upvote;
    }

    public Integer getDownvote() {
        return downvote;
    }

    public void setDownvote(Integer downvote) {
        this.downvote = downvote;
    }

    public Integer getSeen() {
        return seen;
    }

    public void setSeen(Integer seen) {
        this.seen = seen;
    }

    public Integer getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(Integer createdAt) {
        this.createdAt = createdAt;
    }

    public String getCreatorFullName() {
        return creatorFullName;
    }

    public void setCreatorFullName(String creatorFullName) {
        this.creatorFullName = creatorFullName;
    }

    public String getCreatorEmail() {
        return creatorEmail;
    }

    public void setCreatorEmail(String creatorEmail) {
        this.creatorEmail = creatorEmail;
    }

    public String getMountainName() {
        return mountainName;
    }

    public void setMountainName(String mountainName) {
        this.mountainName = mountainName;
    }

    public String getMountainPhoto() {
        return mountainPhoto;
    }

    public void setMountainPhoto(String mountainPhoto) {
        this.mountainPhoto = mountainPhoto;
    }
}
