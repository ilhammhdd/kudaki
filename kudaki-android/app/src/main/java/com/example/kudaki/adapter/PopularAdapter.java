package com.example.kudaki.adapter;

import android.content.Context;
import android.content.Intent;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.example.kudaki.R;
import com.example.kudaki.explore.MountainActivity;
import com.example.kudaki.model.response.Mountain;

import java.util.ArrayList;

import butterknife.BindView;
import butterknife.ButterKnife;

public class PopularAdapter extends RecyclerView.Adapter<PopularAdapter.ViewHolder> {
    private Context context;
    private ArrayList<Mountain> list;

    public PopularAdapter(Context context, ArrayList<Mountain> list) {
        this.context = context;
        this.list = list;
    }

    @NonNull
    @Override
    public ViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext())
                .inflate(R.layout.popular_item, parent, false);
        return new ViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolder holder, int position) {
        holder.name.setText("Gunung " + list.get(position).getName());
        holder.height.setText(list.get(position).getHeight() + " Mdpl");

        Glide.with(context)
                .load(list.get(position).getPhotos().get(0).getFilePath())
                .into(holder.image);

        holder.btnDetail.setOnClickListener(v -> {
            Intent intent = new Intent(v.getContext(), MountainActivity.class);
            intent.putExtra("uuid", list.get(position).getUuid());
            intent.putExtra("name", "Gunung " + list.get(position).getName());
            intent.putExtra("photo", list.get(position).getPhotos().get(0).getFilePath());
            intent.putExtra("height", list.get(position).getHeight());
            intent.putExtra("latitude", list.get(position).getLatitude());
            intent.putExtra("longitude", list.get(position).getLongitude());
            intent.putExtra("difficulty", list.get(position).getDifficulty());
            intent.putExtra("description", list.get(position).getDescription());
            context.startActivity(intent);
        });
    }

    @Override
    public int getItemCount() {
        return list.size();
    }

    class ViewHolder extends RecyclerView.ViewHolder {
        @BindView(R.id.popularDetail)
        Button btnDetail;
        @BindView(R.id.popularImage)
        ImageView image;
        @BindView(R.id.popularName)
        TextView name;
        @BindView(R.id.popularHeight)
        TextView height;

        ViewHolder(@NonNull View itemView) {
            super(itemView);
            ButterKnife.bind(this, itemView);
        }
    }
}
