package com.example.kudaki.renting;

import android.app.ProgressDialog;
import android.content.Intent;
import android.os.Bundle;
import android.util.Log;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.Button;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.constraintlayout.widget.ConstraintLayout;
import androidx.core.app.NavUtils;

import com.bumptech.glide.Glide;
import com.example.kudaki.R;
import com.example.kudaki.cart.CartActivity;
import com.example.kudaki.transaction.TenantTransactionActivity;
import com.orhanobut.hawk.Hawk;

import java.text.NumberFormat;
import java.util.Locale;

import butterknife.BindView;
import butterknife.ButterKnife;

public class DetailEquipmentActivity extends AppCompatActivity implements DetailEquipmentContract.View {
    @BindView(R.id.detailEquipmentToolbar)
    Toolbar toolbar;
    @BindView(R.id.detailEquipmentImage)
    ImageView image;
    @BindView(R.id.detailEquipmentName)
    TextView name;
    @BindView(R.id.detailEquipmentRating)
    TextView rating;
    @BindView(R.id.detailEquipmentPrice)
    TextView price;
    @BindView(R.id.detailEquipmentAmount)
    TextView amount;
    @BindView(R.id.detailEquipmentDesc)
    TextView desc;
    @BindView(R.id.detailEquipmentAdd)
    ConstraintLayout add;
    @BindView(R.id.detailEquipmentCheckout)
    Button checkout;

    String token, uuid;

    ProgressDialog progressDialog;

    DetailEquipmentContract.Presenter contractPresenter;
    DetailEquipmentPresenter presenter;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_detail_equipment);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        Hawk.init(this).build();

        token = Hawk.get("token");

        uuid = getIntent().getStringExtra("uuid");

        progressDialog = new ProgressDialog(this);

        presenter = new DetailEquipmentPresenter(this, token, uuid);
    }

    @Override
    protected void onStart() {
        super.onStart();

        Locale localeID = new Locale("in", "ID");
        NumberFormat formatRupiah = NumberFormat.getCurrencyInstance(localeID);

        Glide.with(this)
                .load("https://www.static-src.com/wcsstore/Indraprastha/images/catalog/medium//760/eiger_eiger-tas-daypack-base-camp---hitam_full04.jpg")
                .into(image);

        Log.d("PHOTO", "onStart: " + getIntent().getStringExtra("photo"));

        name.setText(getIntent().getStringExtra("name"));
        amount.setText("Stok: " + getIntent().getIntExtra("amount", 0));
        price.setText(formatRupiah.format(getIntent().getIntExtra("price", 0)));
        desc.setText(getIntent().getStringExtra("desc"));
        rating.setText(String.valueOf(getIntent().getDoubleExtra("rating", 0)));
    }

    @Override
    protected void onResume() {
        super.onResume();

        add.setOnClickListener(v -> contractPresenter.addItem());

        checkout.setOnClickListener(v -> contractPresenter.checkout());
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        getMenuInflater().inflate(R.menu.renting_menu, menu);
        return super.onCreateOptionsMenu(menu);
    }

    @Override
    public boolean onOptionsItemSelected(@NonNull MenuItem item) {
        switch (item.getItemId()) {
            case R.id.shopping_cart:
                Intent cart = new Intent(this, CartActivity.class);
                startActivity(cart);
                return true;
            case android.R.id.home:
                NavUtils.navigateUpFromSameTask(this);
                return true;
        }

        return super.onOptionsItemSelected(item);
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
    public void showAddSuccess(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
    }

    @Override
    public void showCheckoutSuccess(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
        Intent transtaction = new Intent(this, TenantTransactionActivity.class);
        transtaction.putExtra("status", "PENDING");
        startActivity(transtaction);
    }

    @Override
    public void setPresenter(DetailEquipmentContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }
}
