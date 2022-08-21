package com.example.kudaki.adapter;

import android.content.Context;
import android.content.Intent;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AlertDialog;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.kudaki.R;
import com.example.kudaki.explore.AddRecommendationItemActivity;
import com.example.kudaki.model.response.RecommendationData;
import com.example.kudaki.model.response.RecommendationResponse;
import com.example.kudaki.model.response.RecommendedGear;
import com.example.kudaki.retrofit.GetData;
import com.example.kudaki.retrofit.RetrofitClient;

import java.util.ArrayList;

import butterknife.BindView;
import butterknife.ButterKnife;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class RecommendationAdapter extends RecyclerView.Adapter<RecommendationAdapter.ViewHolder> {
    Context context;
    ArrayList<RecommendedGear> list;

    String token;

    public RecommendationAdapter(Context context, ArrayList<RecommendedGear> list) {
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
                .inflate(R.layout.recommendation_item, parent, false);
        return new ViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolder holder, int position) {
        Log.d("NAME", "onBindViewHolder: " + list.get(position).getCreatorFullName());
        holder.name.setText(list.get(position).getCreatorFullName());
        holder.seen.setText(String.valueOf(list.get(position).getSeen()));
        holder.detail.setOnClickListener(v -> {
            String uuid = list.get(position).getUuid();

            AlertDialog.Builder builder = new AlertDialog.Builder(context, R.style.CustomDialogTheme);

            LayoutInflater inflater = LayoutInflater.from(context);

            View view = inflater.inflate(R.layout.dialog_recommend_item, null);
            RecyclerView recyclerView = view.findViewById(R.id.rvRecommendItem);

            builder.setTitle("Rekomendasi Alat");
            builder.setNeutralButton("Tutup", (dialog, which) -> dialog.dismiss());
            builder.setPositiveButton("Tambah Alat Lain", (dialog, which) -> {
                Intent intent = new Intent(context, AddRecommendationItemActivity.class);
                intent.putExtra("uuid", list.get(position).getUuid());
                context.startActivity(intent);
            });
            builder.setView(view);

            GetData service = RetrofitClient.getRetrofit().create(GetData.class);
            Call<RecommendationResponse> call = service.getRecommendationItems(token, uuid, 10, 0);

            call.enqueue(new Callback<RecommendationResponse>() {
                @Override
                public void onResponse(Call<RecommendationResponse> call, Response<RecommendationResponse> response) {
                    if (response.code() == 200) {
                        RecommendationResponse resp = response.body();

                        RecommendationData data = resp.getData();

                        RecommendDetailAdapter adapter = new RecommendDetailAdapter(context, data.getRecommendedGearItems());
                        recyclerView.setAdapter(adapter);
                        recyclerView.setLayoutManager(new LinearLayoutManager(context));
                    }
                }

                @Override
                public void onFailure(Call<RecommendationResponse> call, Throwable t) {

                }
            });
            builder.show();
        });
    }

    @Override
    public int getItemCount() {
        return list.size();
    }

    public class ViewHolder extends RecyclerView.ViewHolder {
        @BindView(R.id.recommendName)
        TextView name;
        @BindView(R.id.recommendSeen)
        TextView seen;
        @BindView(R.id.recommendDetail)
        Button detail;
        public ViewHolder(@NonNull View itemView) {
            super(itemView);
            ButterKnife.bind(this, itemView);
        }
    }
}
