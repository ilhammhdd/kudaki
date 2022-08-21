package com.example.kudaki.transaction;

import android.app.ProgressDialog;
import android.os.Bundle;
import android.util.Log;
import android.view.MenuItem;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.core.app.NavUtils;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.kudaki.R;
import com.example.kudaki.adapter.TenantTransactionAdapter;
import com.example.kudaki.model.response.Order;
import com.example.kudaki.model.response.OrderHistoryData;
import com.orhanobut.hawk.Hawk;

import java.util.ArrayList;

import butterknife.BindView;
import butterknife.ButterKnife;

public class TenantTransactionActivity extends AppCompatActivity implements TenantTransactionContract.View {
    @BindView(R.id.transactionToolbar)
    Toolbar toolbar;
    @BindView(R.id.rvTransaction)
    RecyclerView recyclerView;

    String token, status;
    ArrayList<Order> list;
    TenantTransactionAdapter adapter;

    TenantTransactionContract.Presenter contractPresenter;
    TenantTransactionPresenter presenter;

    ProgressDialog progressDialog;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_transaction);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        Hawk.init(this).build();

        token = Hawk.get("token");

        Log.d("status", "onCreate: " + getIntent());
        status = "PENDING";
        if (!getIntent().getStringExtra("status").isEmpty()) {
            status = getIntent().getExtras().getString("status");
        }

        list = new ArrayList<>();

        presenter = new TenantTransactionPresenter(this, token);

        progressDialog = new ProgressDialog(this);
    }

    @Override
    protected void onStart() {
        super.onStart();

        contractPresenter.loadTransaction(status);
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        switch (item.getItemId()) {
            case android.R.id.home:
                NavUtils.navigateUpFromSameTask(this);
                return true;
        }
        return super.onOptionsItemSelected(item);
    }

    @Override
    public void setPresenter(TenantTransactionContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }

    @Override
    public void showProgress() {
        progressDialog.setMax(100);
        progressDialog.setMessage("Please wait...");
        progressDialog.setTitle("Loading");
        progressDialog.setProgressStyle(ProgressDialog.STYLE_SPINNER);
        progressDialog.show();
    }

    @Override
    public void closeProgress() {
        progressDialog.dismiss();
    }

    @Override
    public void showOrderHistoryData(OrderHistoryData data) {
        if (data.getOrders() == null) {
            Toast.makeText(this, "Transaksi Kosong", Toast.LENGTH_SHORT).show();
        } else {
            list.clear();
            for (int i = 0; i < data.getOrders().size(); i++) {
                list.add(new Order(
                        data.getOrders().get(i).getOrderNum(),
                        data.getOrders().get(i).getStatus(),
                        data.getOrders().get(i).getCreatedAt(),
                        data.getOrders().get(i).getTotalItem(),
                        data.getOrders().get(i).getTotalPrice(),
                        data.getOrders().get(i).getOwners()
                ));
            }
            adapter = new TenantTransactionAdapter(this, data.getOrders());
            adapter.notifyDataSetChanged();
            adapter.setToken(token);
            recyclerView.setLayoutManager(new LinearLayoutManager(this, RecyclerView.VERTICAL, false));
            recyclerView.setAdapter(adapter);
        }
    }
}
