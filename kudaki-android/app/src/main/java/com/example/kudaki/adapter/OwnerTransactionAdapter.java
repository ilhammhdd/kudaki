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
import com.example.kudaki.model.response.OrderOwner;
import com.example.kudaki.transaction.OwnerDetailTransactionActivity;

import java.text.NumberFormat;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.Locale;
import java.util.TimeZone;

import butterknife.BindView;
import butterknife.ButterKnife;

public class OwnerTransactionAdapter extends RecyclerView.Adapter<OwnerTransactionAdapter.ViewHolder> {
    Context context;
    ArrayList<OrderOwner> list;

    String token;

    public OwnerTransactionAdapter(Context context, ArrayList<OrderOwner> list) {
        this.context = context;
        this.list = list;
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
            Locale localeID = new Locale("in", "ID");
            NumberFormat formatRupiah = NumberFormat.getCurrencyInstance(localeID);

            Intent intent = new Intent(context, OwnerDetailTransactionActivity.class);
            intent.putExtra("uuid", list.get(position).getUuid());
            intent.putExtra("name", list.get(position).getTenant().getFullName());
            intent.putExtra("amount", list.get(position).getTotalItem());
            intent.putExtra("price", formatRupiah.format(list.get(position).getTotalPrice()));
            intent.putExtra("status", list.get(position).getStatus());
            context.startActivity(intent);
        });
    }

    @Override
    public int getItemCount() {
        return list.size();
    }

    public void setToken(String token) {
        this.token = token;
    }

    public class ViewHolder extends RecyclerView.ViewHolder {
        @BindView(R.id.transItemDate)
        TextView date;
        @BindView(R.id.transItemStatus)
        TextView status;
        @BindView(R.id.transactionItem)
        CardView cardView;
        public ViewHolder(@NonNull View itemView) {
            super(itemView);
            ButterKnife.bind(this, itemView);
        }
    }
}
