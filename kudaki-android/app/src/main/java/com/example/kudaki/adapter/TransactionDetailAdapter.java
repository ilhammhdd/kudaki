package com.example.kudaki.adapter;

import android.content.Context;
import android.content.Intent;
import android.net.Uri;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.EditText;
import android.widget.RadioGroup;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AlertDialog;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.kudaki.R;
import com.example.kudaki.model.response.DefaultResponse;
import com.example.kudaki.model.response.Owner;
import com.example.kudaki.retrofit.PostData;
import com.example.kudaki.retrofit.RetrofitClient;

import java.util.ArrayList;

import butterknife.BindView;
import butterknife.ButterKnife;
import okhttp3.MultipartBody;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class TransactionDetailAdapter extends RecyclerView.Adapter<TransactionDetailAdapter.ViewHolder> {
    Context context;
    ArrayList<Owner> list;

    String token;

    public TransactionDetailAdapter(Context context, ArrayList<Owner> list) {
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
                .inflate(R.layout.transaction_detail_item, parent, false);
        return new ViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolder holder, int position) {
        if (list.get(position).getOwnerApprovalStatus().equalsIgnoreCase("DONE")) {
            Log.d("DETAIL", "onBindViewHolder: " + true);

            AlertDialog.Builder builder = new AlertDialog.Builder(context, R.style.CustomDialogTheme);

            builder.setTitle("Penyewaan selesai");
            builder.setMessage("Silahkan berikan penilaianmu untuk " + list.get(position).getFullName() +
                    " dengan menekan tombol Review di halaman ini!");
            builder.setCancelable(true);
            builder.setNeutralButton("Tutup", (dialog, which) -> dialog.dismiss());
            builder.show();

            holder.review.setVisibility(View.VISIBLE);
            holder.review.setOnClickListener(v -> {
                AlertDialog.Builder builder1 = new AlertDialog.Builder(context, R.style.CustomDialogTheme);

                LayoutInflater inflater = LayoutInflater.from(context);

                View view = inflater.inflate(R.layout.dialog_review, null);
                EditText review = view.findViewById(R.id.orderReview);
                RadioGroup rating = view.findViewById(R.id.orderRating);

                rating.setOnCheckedChangeListener(new RadioGroup.OnCheckedChangeListener() {
                    @Override
                    public void onCheckedChanged(RadioGroup group, int checkedId) {

                    }
                });

                builder1.setTitle("Berikan Penilaianmu");
                builder1.setView(view);
                builder1.setPositiveButton("Kirim", (dialog, which) -> {
                    if (review.getText().toString().isEmpty()) {
                        Toast.makeText(v.getContext(), "Mohon berikan penilaianmu terlebih dahulu", Toast.LENGTH_SHORT).show();
                    } else {
                        PostData service = RetrofitClient.getRetrofit().create(PostData.class);
                        RequestBody requestBody = new MultipartBody.Builder()
                                .setType(MultipartBody.FORM)
                                .addFormDataPart("rating", "")
                                .addFormDataPart("owner_order_uuid", list.get(position).getOwnerOrderUuid())
                                .addFormDataPart("review", review.getText().toString())
                                .build();
                        Call<DefaultResponse> call = service.addReview(token, requestBody);

                        call.enqueue(new Callback<DefaultResponse>() {
                            @Override
                            public void onResponse(Call<DefaultResponse> call, Response<DefaultResponse> response) {
                                if (response.code() == 200) {
                                    Toast.makeText(v.getContext(), "Terima kasih atas penilaianmu", Toast.LENGTH_SHORT).show();
                                }

                                dialog.dismiss();
                            }

                            @Override
                            public void onFailure(Call<DefaultResponse> call, Throwable t) {

                            }
                        });
                    }
                });

                builder1.show();
            });
        }
        holder.name.setText(list.get(position).getFullName());
        holder.phone.setOnClickListener(v -> {
            try {
                String text = list.get(position).getPhoneNumber();

                String toNumber = text.replaceFirst("0", "62");

                Log.d("Phone", "onClick: " + toNumber);

                Intent intent = new Intent(Intent.ACTION_VIEW);
                intent.setData(Uri.parse("http://api.whatsapp.com/send?phone="+toNumber));
                context.startActivity(intent);
            }
            catch (Exception e){
                e.printStackTrace();
            }
        });

        TransactionDetailItemAdapter adapter = new TransactionDetailItemAdapter(context, list.get(position).getCartItems());
        holder.recyclerView.setLayoutManager(new LinearLayoutManager(context, RecyclerView.VERTICAL, false));
        holder.recyclerView.setAdapter(adapter);
    }

    @Override
    public int getItemCount() {
        return list.size();
    }

    public class ViewHolder extends RecyclerView.ViewHolder {
        @BindView(R.id.transactionDetailName)
        TextView name;
        @BindView(R.id.transactionDetailPhone)
        TextView phone;
        @BindView(R.id.transactionDetailReview)
        Button review;
        @BindView(R.id.rvTransactionDetail)
        RecyclerView recyclerView;

        public ViewHolder(@NonNull View itemView) {
            super(itemView);
            ButterKnife.bind(this, itemView);
        }
    }
}
