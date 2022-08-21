package com.example.kudaki.transaction;

import android.os.Bundle;
import android.view.MenuItem;
import android.widget.TextView;

import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.core.app.NavUtils;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.kudaki.R;
import com.example.kudaki.adapter.TransactionDetailAdapter;
import com.example.kudaki.model.response.Owner;
import com.orhanobut.hawk.Hawk;

import java.util.ArrayList;

import butterknife.BindView;
import butterknife.ButterKnife;

public class TenantDetailTransactionActivity extends AppCompatActivity {
    @BindView(R.id.tenantTransactionToolbar)
    Toolbar toolbar;
    @BindView(R.id.rvTenantItem)
    RecyclerView recyclerView;
    @BindView(R.id.tenantTransactionAmount)
    TextView amount;
    @BindView(R.id.tenantTransactionPrice)
    TextView price;

    String token;
    ArrayList<Owner> list;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_tenant_detail_transaction);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        Hawk.init(this).build();

        token = Hawk.get("token");
        list = Hawk.get("owners");

        TransactionDetailAdapter adapter = new TransactionDetailAdapter(this, list);
        adapter.setToken(token);
        recyclerView.setLayoutManager(new LinearLayoutManager(this));
        recyclerView.setAdapter(adapter);

        amount.setText(getIntent().getExtras().getString("amount"));
        price.setText(getIntent().getExtras().getString("price"));
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        switch (item.getItemId()) {
            case android.R.id.home:
                Hawk.delete("owners");
                NavUtils.navigateUpFromSameTask(this);
                return true;
        }
        return super.onOptionsItemSelected(item);
    }
}
