package com.example.kudaki.retrofit;

import com.example.kudaki.model.response.AddressResponse;
import com.example.kudaki.model.response.AllItemResponse;
import com.example.kudaki.model.response.FileResponse;
import com.example.kudaki.model.response.RecommendationResponse;
import com.example.kudaki.model.response.CartResponse;
import com.example.kudaki.model.response.MountainResponse;
import com.example.kudaki.model.response.OrderHistoryResponse;
import com.example.kudaki.model.response.OwnerHistoryReponse;
import com.example.kudaki.model.response.ProfileResponse;
import com.example.kudaki.model.response.StoreResponse;

import retrofit2.Call;
import retrofit2.http.GET;
import retrofit2.http.Header;
import retrofit2.http.Query;

public interface GetData {
    @GET("/file")
    Call<FileResponse> getFile();

    @GET("/mountain")
    Call<MountainResponse> getAllMountain(@Header("Kudaki-Token") String token,
                                          @Query("limit") int limit,
                                          @Query("offset") int offset);

    @GET("/recommendations")
    Call<RecommendationResponse> getAllRecommendation(@Header("Kudaki-Token") String token,
                                                      @Query("mountain_uuid") String uuid,
                                                      @Query("limit") int limit,
                                                      @Query("offset") int offset);

    @GET("/recommendation/items")
    Call<RecommendationResponse> getRecommendationItems(@Header("Kudaki-Token") String token,
                                                      @Query("recommended_gear_uuid") String uuid,
                                                      @Query("limit") int limit,
                                                      @Query("offset") int offset);

    @GET("/user-info/profile")
    Call<ProfileResponse> getProfile(@Header("Kudaki-Token") String token);

    @GET("/user-info/addresses")
    Call<AddressResponse> getAddress(@Header("Kudaki-Token") String token);

    @GET("/storefront/items")
    Call<StoreResponse> getStoreItems(@Header ("Kudaki-Token") String token,
                                      @Query("limit") int limit,
                                      @Query("offset") int offset);

    @GET("/rental/cart/items")
    Call<CartResponse> getCartItems(@Header ("Kudaki-Token") String token,
                                    @Query("offset") int offset,
                                    @Query("limit") int limit);

    @GET("/items")
    Call<AllItemResponse> getAllItems(@Header ("Kudaki-Token") String token,
                                      @Query("offset") int offset,
                                      @Query("limit") int limit);

    @GET("/item/search")
    Call<AllItemResponse> searchItems(@Header ("Kudaki-Token") String token,
                                     @Query("keyword") String keyword,
                                     @Query("offset") int offset,
                                     @Query("limit") int limit);

    @GET("/order/owner")
    Call<OwnerHistoryReponse> ownerOrderHistory(@Header ("Kudaki-Token") String token,
                                                @Query("limit") int limit,
                                                @Query("offset") int offset,
                                                @Query("order_status") String status);

    @GET("/order/tenant")
    Call<OrderHistoryResponse> getOrderHistory(@Header ("Kudaki-Token") String token,
                                               @Query("limit") int limit,
                                               @Query("offset") int offset,
                                               @Query("order_status") String status);

}
