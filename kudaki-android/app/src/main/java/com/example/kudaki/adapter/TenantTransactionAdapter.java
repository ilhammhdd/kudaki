package com.example.kudaki.adapter;

import android.content.Context;
import android.content.Intent;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.cardview.widget.CardView;
import androidx.recyclerview.widget.RecyclerView;

import com.example.kudaki.R;
import com.example.kudaki.model.response.Order;
import com.example.kudaki.transaction.TenantDetailTransactionActivity;
import com.orhanobut.hawk.Hawk;

import java.text.NumberFormat;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.Locale;
import java.util.TimeZone;

import butterknife.BindView;
import butterknife.ButterKnife;

public class TenantTransactionAdapter extends RecyclerView.Adapter<TenantTransactionAdapter.ViewHolder> {
    Context context;
    ArrayList<Order> list;

    String token;

    public TenantTransactionAdapter(Context context, ArrayList<Order> list) {
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
                .inflate(R.layout.transaction_item, parent, false);
        return new ViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolder holder, int position) {
        long epoch = list.get(position).getCreatedAt();
        Date date = new Date(epoch*1000L);
        SimpleDateFormat dateFormat = new SimpleDateFormat("dd/MM/yyyy");
        dateFormat.setTimeZone(TimeZone.getTimeZone("GMT+8"));
        holder.date.setText(dateFormat.format(date));
        holder.status.setText(list.get(position).getStatus());

        holder.cardView.setOnClickListener(v -> {
            Hawk.init(context).build();

            Hawk.put("owners", list.get(position).getOwners());

            Locale localeID = new Locale("in", "ID");
            NumberFormat formatRupiah = NumberFormat.getCurrencyInstance(localeID);

            Intent intent = new Intent(context, TenantDetailTransactionActivity.class);
            intent.putExtra("amount", String.valueOf(list.get(position).getTotalItem()));
            intent.putExtra("price", formatRupiah.format(list.get(position).getTotalPrice()));
            context.startActivity(intent);
        });
    }

    @Override
    public int getItemCount() {
        return list.size();
    }

    public class ViewHolder extends RecyclerView.ViewHolder {
        @BindView(R.id.transactionItem)
        CardView cardView;
        @BindView(R.id.transItemDate)
        TextView date;
        @BindView(R.id.transItemStatus)
        TextView status;

        public ViewHolder(@NonNull View itemView) {
            super(itemView);
            ButterKnife.bind(this, itemView);
        }
    }
}
