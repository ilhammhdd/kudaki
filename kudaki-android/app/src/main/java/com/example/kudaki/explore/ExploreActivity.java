package com.example.kudaki.explore;

import android.app.ProgressDialog;
import android.content.Intent;
import android.os.Bundle;

import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.kudaki.MainActivity;
import com.example.kudaki.R;
import com.example.kudaki.adapter.MountainAdapter;
import com.example.kudaki.event.EventActivity;
import com.example.kudaki.model.response.MountainData;
import com.example.kudaki.profile.ProfileActivity;
import com.example.kudaki.renting.RentalActivity;
import com.google.android.material.bottomnavigation.BottomNavigationView;
import com.orhanobut.hawk.Hawk;

import butterknife.BindView;
import butterknife.ButterKnife;

public class ExploreActivity extends AppCompatActivity implements ExploreContract.View {
    @BindView(R.id.exploreNav)
    BottomNavigationView bottomNav;
    @BindView(R.id.rvMountain)
    RecyclerView recyclerView;
    @BindView(R.id.exploreToolbar)
    Toolbar toolbar;

    String token;

    ExploreContract.Presenter contractPresenter;
    ExplorePresenter presenter;

    MountainAdapter adapter;

    ProgressDialog progressDialog;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_explore);
        overridePendingTransition(android.R.anim.fade_in, android.R.anim.fade_out);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);

        Hawk.init(this).build();

        token = Hawk.get("token");

        progressDialog = new ProgressDialog(this);

        presenter = new ExplorePresenter(this, token);
    }

    @Override
    protected void onStart() {
        super.onStart();

        contractPresenter.loadMountain();
    }

    @Override
    protected void onResume() {
        super.onResume();

        bottomNav.getMenu().getItem(2).setChecked(true);
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
                    return true;
                case R.id.navRental:
                    startActivity(new Intent(this, RentalActivity.class));
                    finish();
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
    public void showMountainData(MountainData data) {
        if (data.getMountains() == null) {

        } else {
            adapter = new MountainAdapter(this, data.getMountains());
            adapter.notifyDataSetChanged();
            recyclerView.setLayoutManager(new LinearLayoutManager(this, RecyclerView.VERTICAL, false));
            recyclerView.setAdapter(adapter);
        }
    }

    @Override
    public void setPresenter(ExploreContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }
}
