package com.example.kudaki.renting;

import android.app.ProgressDialog;
import android.content.Intent;
import android.os.Bundle;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.recyclerview.widget.GridLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.kudaki.MainActivity;
import com.example.kudaki.R;
import com.example.kudaki.adapter.RentalAdapter;
import com.example.kudaki.cart.CartActivity;
import com.example.kudaki.event.EventActivity;
import com.example.kudaki.explore.ExploreActivity;
import com.example.kudaki.model.response.AllItemData;
import com.example.kudaki.model.response.StoreItem;
import com.example.kudaki.profile.ProfileActivity;
import com.google.android.material.bottomnavigation.BottomNavigationView;
import com.orhanobut.hawk.Hawk;

import java.util.ArrayList;

import butterknife.BindView;
import butterknife.ButterKnife;

public class RentalActivity extends AppCompatActivity implements RentalContract.View {
    @BindView(R.id.rentalToolbar)
    Toolbar toolbar;
    @BindView(R.id.rentalNav)
    BottomNavigationView bottomNav;
    @BindView(R.id.rvEquipment)
    RecyclerView recyclerView;

    String token;
    ArrayList<StoreItem> list;
    RentalAdapter adapter;

    RentalContract.Presenter contractPresenter;
    RentalPresenter presenter;

    ProgressDialog progressDialog;

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
        }

        return super.onOptionsItemSelected(item);
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_rental);
        overridePendingTransition(android.R.anim.fade_in, android.R.anim.fade_out);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);

        Hawk.init(this).build();

        token = Hawk.get("token");

        progressDialog = new ProgressDialog(this);

        presenter = new RentalPresenter(this, token);

        list = new ArrayList<>();
    }

    @Override
    protected void onStart() {
        super.onStart();

        contractPresenter.loadItems();
    }

    @Override
    protected void onResume() {
        super.onResume();

        bottomNav.getMenu().getItem(3).setChecked(true);
        bottomNav.setOnNavigationItemSelectedListener(menuItem -> {
            switch (menuItem.getItemId()) {
                case R.id.navHome:
                    startActivity(new Intent(this, MainActivity.class));
                    finish();
                    return true;
                case R.id.navEvent:
                    startActivity(new Intent(this, EventActivity.class));
                    finish();
                    return true;
                case R.id.navExplore:
                    startActivity(new Intent(this, ExploreActivity.class));
                    finish();
                    return true;
                case R.id.navRental:
                    return true;
                case R.id.navProfile:
                    startActivity(new Intent(this, ProfileActivity.class));
                    finish();
                    return true;
            }
            return false;
        });
    }

    @Override
    public void setPresenter(RentalContract.Presenter presenter) {
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
    public void displayItems(AllItemData data) {
        if (data.getItems() == null) {
            Toast.makeText(this, "List Alat Kosong", Toast.LENGTH_SHORT).show();
        } else {
            list.clear();
            for (int i = 0; i < data.getItems().size(); i++) {
                list.add(new StoreItem(
                        data.getItems().get(i).getUuid(),
                        data.getItems().get(i).getStorefrontUuid(),
                        data.getItems().get(i).getName(),
                        data.getItems().get(i).getPrice(),
                        data.getItems().get(i).getPhoto(),
                        data.getItems().get(i).getDescription(),
                        data.getItems().get(i).getRating(),
                        data.getItems().get(i).getPriceDuration()
                ));
            }
        }
        adapter = new RentalAdapter(this, data.getItems());
        adapter.notifyDataSetChanged();
        adapter.setToken(token);
        recyclerView.setLayoutManager(new GridLayoutManager(this, 2));
        recyclerView.setAdapter(adapter);
    }
}
