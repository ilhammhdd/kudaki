package com.example.kudaki.transaction;

import android.app.ProgressDialog;
import android.os.Bundle;
import android.view.MenuItem;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.core.app.NavUtils;

import com.example.kudaki.R;
import com.orhanobut.hawk.Hawk;

import butterknife.BindView;
import butterknife.ButterKnife;

public class OwnerDetailTransactionActivity extends AppCompatActivity implements OwnerDetailTransactionContract.View {
    @BindView(R.id.ownerTransactionToolbar)
    Toolbar toolbar;
    @BindView(R.id.ownerDetailName)
    TextView name;
    @BindView(R.id.ownerDetailAmount)
    TextView amount;
    @BindView(R.id.ownerDetailPrice)
    TextView price;
    @BindView(R.id.ownerDetailAction)
    Button button;
    @BindView(R.id.ownerDetailActionSecondary)
    Button buttonSecondary;

    String token, status, uuid;

    OwnerDetailTransactionContract.Presenter contractPresenter;
    OwnerDetailTransactionPresenter presenter;

    ProgressDialog progressDialog;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_owner_detail_transaction);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        Hawk.init(this).build();

        token = Hawk.get("token");
        status = getIntent().getExtras().getString("status");
        uuid = getIntent().getExtras().getString("uuid");

        presenter = new OwnerDetailTransactionPresenter(this, token, uuid);

        progressDialog = new ProgressDialog(this);

        name.setText(getIntent().getExtras().getString("name"));
        amount.setText(String.valueOf(getIntent().getExtras().getInt("amount")));
        price.setText(getIntent().getExtras().getString("price"));
    }

    @Override
    protected void onResume() {
        super.onResume();

        switch (status) {
            case "PENDING":
                button.setText("Terima Permintaan Sewa");
                buttonSecondary.setText("Tolak Permintaan Sewa");

                button.setOnClickListener(v -> contractPresenter.approve());

                buttonSecondary.setOnClickListener(v -> contractPresenter.disapprove());
                break;
            case "APPROVED":
                buttonSecondary.setVisibility(View.GONE);
                button.setText("Konfirmasi Barang Disewa");

                button.setOnClickListener(v -> contractPresenter.rented());
                break;
            case "RENTED":
                buttonSecondary.setVisibility(View.GONE);
                button.setText("Konfirmasi Sewa Selesai");

                button.setOnClickListener(v -> contractPresenter.done());
                break;
            default:
                button.setVisibility(View.GONE);
                buttonSecondary.setVisibility(View.GONE);
                break;
        }
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
    public boolean onOptionsItemSelected(MenuItem item) {
        switch (item.getItemId()) {
            case android.R.id.home:
                Hawk.delete("owners");
                NavUtils.navigateUpFromSameTask(this);
                return true;
        }
        return super.onOptionsItemSelected(item);
    }

    @Override
    public void showActionSuccess(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
        NavUtils.navigateUpFromSameTask(this);
    }

    @Override
    public void setPresenter(OwnerDetailTransactionContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }
}
