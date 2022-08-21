package com.example.kudaki.retrofit;

import com.example.kudaki.model.request.RecommendationRequest;
import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.model.response.FileResponse;
import com.example.kudaki.model.response.LoginResponse;

import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.DELETE;
import retrofit2.http.Header;
import retrofit2.http.PATCH;
import retrofit2.http.POST;
import retrofit2.http.PUT;
import retrofit2.http.Query;

public interface PostData {
    @POST("/order/checkout")
    Call<DefaultResponse> checkout(@Header ("Kudaki-Token") String token,
                                   @Body RequestBody user);

    @POST("/recommendation")
    Call<DefaultResponse> addRecommendation(@Header ("Kudaki-Token") String token,
                                            @Header ("Content-Type") String contentType,
                                            @Body RecommendationRequest body);

    @POST("/recommendation/item")
    Call<DefaultResponse> addRecommendationItem(@Header ("Kudaki-Token") String token,
                                            @Body RequestBody body);

    @POST("/file")
    Call<FileResponse> uploadFile(@Body RequestBody user);

    @POST("/order/owner-order-review")
    Call<DefaultResponse> addReview(@Header ("Kudaki-Token") String token,
                                   @Body RequestBody user);

    @POST("/login")
    Call<LoginResponse> loginUser(@Body RequestBody user);

    @POST("/signup")
    Call<DefaultResponse> registerUser(@Body RequestBody user);

    @POST("/user-info/address")
    Call<DefaultResponse> addAddress(@Header ("Kudaki-Token") String token,
                                     @Body RequestBody user);

    @POST("/user/password/reset")
    Call<DefaultResponse> sendForgotPwdEmail(@Body RequestBody user);

    @PATCH("/user/password/change")
    Call<DefaultResponse> changePwd(@Header ("Kudaki-Token") String token,
                                    @Body RequestBody user);

    @PATCH("/user/password/reset")
    Call<DefaultResponse> resetPwd(@Query("reset_token") String token,
                                   @Body RequestBody user);

    @POST("/rental/cart/item")
    Call<DefaultResponse> addToCart(@Header ("Kudaki-Token") String token,
                                    @Body RequestBody user);

    @DELETE("/rental/cart/item")
    Call<DefaultResponse> deleteCartItem(@Header ("Kudaki-Token") String token,
                                         @Query("cart_item_uuid") String uuid);

    @PATCH("/rental/cart/item")
    Call<DefaultResponse> updateCartItem(@Header ("Kudaki-Token") String token,
                                         @Body RequestBody user);

    @PATCH("/user-info/profile")
    Call<DefaultResponse> updateProfile(@Header ("Kudaki-Token") String token,
                                        @Body RequestBody user);

    @POST("/storefront/item")
    Call<DefaultResponse> addStoreItem(@Header ("Kudaki-Token") String token,
                                       @Body RequestBody user);

    @DELETE("/storefront/item")
    Call<DefaultResponse> deleteStoreItem(@Header ("Kudaki-Token") String token,
                                          @Query("item_uuid") String uuid);

    @PUT("/storefront/item")
    Call<DefaultResponse> updateStoreItem(@Header ("Kudaki-Token") String token,
                                          @Body RequestBody user);

    @POST("/order/owner/rented")
    Call<DefaultResponse> confirmRented(@Header ("Kudaki-Token") String token,
                                       @Body RequestBody user);

    @POST("/order/owner/approve")
    Call<DefaultResponse> approveOrder(@Header ("Kudaki-Token") String token,
                                       @Body RequestBody user);

    @POST("/order/confirm-returnment/owner")
    Call<DefaultResponse> confirmReturn(@Header ("Kudaki-Token") String token,
                                       @Body RequestBody user);

    @POST("/order/owner/disapprove")
    Call<DefaultResponse> dissaproveOrder(@Header ("Kudaki-Token") String token,
                                          @Body RequestBody user);
}
