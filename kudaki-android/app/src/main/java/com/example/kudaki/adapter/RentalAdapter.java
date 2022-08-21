package com.example.kudaki.adapter;

import android.app.ProgressDialog;
import android.content.Context;
import android.content.Intent;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.cardview.widget.CardView;
import androidx.recyclerview.widget.RecyclerView;

import com.example.kudaki.R;
import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.model.response.StoreItem;
import com.example.kudaki.renting.DetailEquipmentActivity;
import com.example.kudaki.retrofit.PostData;
import com.example.kudaki.retrofit.RetrofitClient;
import com.squareup.picasso.Picasso;

import java.text.NumberFormat;
import java.util.ArrayList;
import java.util.Locale;

import butterknife.BindView;
import butterknife.ButterKnife;
import okhttp3.MultipartBody;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class RentalAdapter extends RecyclerView.Adapter<RentalAdapter.ViewHolder> {
    Context context;
    ArrayList<StoreItem> list;

    String token;

    public RentalAdapter(Context context, ArrayList<StoreItem> list) {
        this.context = context;
        this.list = list;
    }

    public void setToken(String token) {
        this.token = token;
    }

    @NonNull
    @Override
    public ViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext())
                .inflate(R.layout.rental_item, parent, false);
        return new RentalAdapter.ViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolder holder, int position) {
        Locale localeID = new Locale("in", "ID");
        NumberFormat formatRupiah = NumberFormat.getCurrencyInstance(localeID);
        holder.title.setText(list.get(position).getName());
        holder.price.setText(formatRupiah.format(list.get(position).getPrice()) + "/hari");
        holder.rating.setText(String.valueOf(list.get(position).getRating()));

        Picasso.get().load("https://www.static-src.com/wcsstore/Indraprastha/images/catalog/medium//760/eiger_eiger-tas-daypack-base-camp---hitam_full04.jpg")
                .into(holder.image);

        holder.btnAdd.setOnClickListener(v -> {
            Long time = System.currentTimeMillis() / 1000L;

            ProgressDialog progressDialog = new ProgressDialog(context);
            progressDialog.setMax(100);
            progressDialog.setMessage("Please wait...");
            progressDialog.setTitle("Loading");
            progressDialog.setProgressStyle(ProgressDialog.STYLE_SPINNER);
            progressDialog.show();

            PostData service = RetrofitClient.getRetrofit().create(PostData.class);
            RequestBody requestBody = new MultipartBody.Builder()
                    .setType(MultipartBody.FORM)
                    .addFormDataPart("item_uuid", list.get(position).getUuid())
                    .addFormDataPart("item_amount", "1")
                    .addFormDataPart("duration_from", String.valueOf(time))
                    .addFormDataPart("duration", "1")
                    .build();

            Call<DefaultResponse> call = service.addToCart(token, requestBody);

            call.enqueue(new Callback<DefaultResponse>() {
                @Override
                public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                    if (response.code() == 200) {
                        Toast.makeText(context, "Berhasil ditambahkan ke keranjang", Toast.LENGTH_SHORT).show();
                    }
                    progressDialog.dismiss();
                }

                @Override
                public void onFailure(Call<DefaultResponse> call, Throwable t) {

                }
            });
        });

        holder.card.setOnClickListener(v -> {
            Intent detail = new Intent(context, DetailEquipmentActivity.class);
            detail.putExtra("uuid", list.get(position).getUuid());
            detail.putExtra("name", list.get(position).getName());
            detail.putExtra("amount", list.get(position).getAmount());
            detail.putExtra("price", list.get(position).getPrice());
            detail.putExtra("desc", list.get(position).getDescription());
            detail.putExtra("photo", list.get(position).getPhoto());
            detail.putExtra("rating", list.get(position).getRating());
            context.startActivity(detail);
        });
    }

    @Override
    public int getItemCount() {
        return list.size();
    }

    class ViewHolder extends RecyclerView.ViewHolder {
        @BindView(R.id.rentalItemCard)
        CardView card;
        @BindView(R.id.rentalItemRating)
        TextView rating;
        @BindView(R.id.rentalItemTitle)
        TextView title;
        @BindView(R.id.rentalItemPrice)
        TextView price;
        @BindView(R.id.rentalItemAdd)
        Button btnAdd;
        @BindView(R.id.rentalItemImage)
        ImageView image;

        ViewHolder(@NonNull View itemView) {
            super(itemView);
            ButterKnife.bind(this, itemView);
        }
    }
}
