package com.example.kudaki.adapter;

import android.content.Context;
import android.content.Intent;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.example.kudaki.R;
import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.model.response.StoreItem;
import com.example.kudaki.profile.etalase.EditEtalaseActivity;
import com.example.kudaki.profile.etalase.EtalaseActivity;
import com.example.kudaki.retrofit.PostData;
import com.example.kudaki.retrofit.RetrofitClient;

import java.text.NumberFormat;
import java.util.ArrayList;
import java.util.Locale;

import butterknife.BindView;
import butterknife.ButterKnife;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class EtalaseAdapter extends RecyclerView.Adapter<EtalaseAdapter.ViewHolder> {
    Context context;
    String token;

    ArrayList<StoreItem> list;

    public EtalaseAdapter(Context context, ArrayList<StoreItem> list) {
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
                .inflate(R.layout.etalase_item, parent, false);
        return new ViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolder holder, int position) {
        Locale localeID = new Locale("in", "ID");
        NumberFormat formatRupiah = NumberFormat.getCurrencyInstance(localeID);
        holder.name.setText(list.get(position).getName());
        holder.priceDuration.setText(formatRupiah.format(list.get(position).getPrice()));

        holder.edit.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent edit = new Intent(v.getContext(), EditEtalaseActivity.class);
                edit.putExtra("uuid", list.get(position).getUuid());
                edit.putExtra("name", list.get(position).getName());
                edit.putExtra("price", list.get(position).getPrice());
                edit.putExtra("description", list.get(position).getDescription());
                edit.putExtra("duration", list.get(position).getPriceDuration());
                v.getContext().startActivity(edit);
            }
        });

        holder.delete.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                PostData service = RetrofitClient.getRetrofit().create(PostData.class);
                Call<DefaultResponse> call = service.deleteStoreItem(token, list.get(position).getUuid());

                call.enqueue(new Callback<DefaultResponse>() {
                    @Override
                    public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                        if (response.code() == 200) {
                            Toast.makeText(context, "Berhasil hapus", Toast.LENGTH_SHORT).show();

                            ((EtalaseActivity) context).finish();
                            context.startActivity(new Intent(v.getContext(), EtalaseActivity.class));
                        } else {
                            Toast.makeText(context, "Gagal hapus", Toast.LENGTH_SHORT).show();
                        }
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
        return list.size();
    }

    public class ViewHolder extends RecyclerView.ViewHolder {
        @BindView(R.id.etalaseItemName)
        TextView name;
        @BindView(R.id.etalaseItemPriceDuration)
        TextView priceDuration;
        @BindView(R.id.etalaseItemEdit)
        ImageView edit;
        @BindView(R.id.etalaseItemDelete)
        ImageView delete;

        public ViewHolder(@NonNull View itemView) {
            super(itemView);
            ButterKnife.bind(this, itemView);
        }
    }
}
