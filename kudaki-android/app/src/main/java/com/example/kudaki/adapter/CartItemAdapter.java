package com.example.kudaki.adapter;

import android.app.ProgressDialog;
import android.content.Context;
import android.content.Intent;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.example.kudaki.R;
import com.example.kudaki.cart.CartActivity;
import com.example.kudaki.model.response.CartItem;
import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.retrofit.PostData;
import com.example.kudaki.retrofit.RetrofitClient;
import com.orhanobut.hawk.Hawk;

import java.text.NumberFormat;
import java.util.ArrayList;
import java.util.Locale;

import butterknife.BindView;
import butterknife.ButterKnife;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class CartItemAdapter extends RecyclerView.Adapter<CartItemAdapter.ViewHolder> {
    private Context context;
    private ArrayList<CartItem> list;

    public CartItemAdapter(Context context, ArrayList<CartItem> list) {
        this.context = context;
        this.list = list;
    }


    @NonNull
    @Override
    public ViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext())
                .inflate(R.layout.cart_equipment_item, parent, false);
        return new ViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolder holder, int position) {
        Locale localeID = new Locale("in", "ID");
        NumberFormat formatRupiah = NumberFormat.getCurrencyInstance(localeID);

        Glide.with(context)
                .load("https://www.static-src.com/wcsstore/Indraprastha/images/catalog/medium//760/eiger_eiger-tas-daypack-base-camp---hitam_full04.jpg")
                .into(holder.imageView);

        holder.name.setText(list.get(position).getItem().getName());
        holder.price.setText(formatRupiah.format(list.get(position).getItem().getPrice()) + "/hari");

        holder.delete.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                ProgressDialog progressDialog = new ProgressDialog(context);
                progressDialog.setMax(100);
                progressDialog.setMessage("Please wait...");
                progressDialog.setTitle("Loading");
                progressDialog.setProgressStyle(ProgressDialog.STYLE_SPINNER);
                progressDialog.show();
                PostData service = RetrofitClient.getRetrofit().create(PostData.class);

                Hawk.init(v.getContext()).build();

                String token = Hawk.get("token");

                Call<DefaultResponse> call = service.deleteCartItem(token, list.get(position).getUuid());

                call.enqueue(new Callback<DefaultResponse>() {
                    @Override
                    public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                        if (response.code() == 200) {
                            Toast.makeText(context, "Berhasil dihapus", Toast.LENGTH_SHORT).show();
                            ((CartActivity) context).finish();
                            context.startActivity(new Intent(context, CartActivity.class));
                        }
                        progressDialog.dismiss();
                    }

                    @Override
                    public void onFailure(Call<DefaultResponse> call, Throwable t) {

                    }
                });
            }
        });
    }

    @Override
    public int getItemCount() {
        Log.d("COUNT", "getItemCount: " + list.size());
        return list.size();
    }

    public class ViewHolder extends RecyclerView.ViewHolder {
        @BindView(R.id.cartName)
        TextView name;
        @BindView(R.id.cartItemPrice)
        TextView price;
        @BindView(R.id.cartItemImage)
        ImageView imageView;
        @BindView(R.id.cartItemDelete)
        TextView delete;

        public ViewHolder(@NonNull View itemView) {
            super(itemView);
            ButterKnife.bind(this, itemView);
        }
    }
}
