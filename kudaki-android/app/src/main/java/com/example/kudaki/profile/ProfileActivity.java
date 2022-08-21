package com.example.kudaki.profile;

import android.app.ProgressDialog;
import android.content.Intent;
import android.os.Bundle;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.viewpager.widget.ViewPager;

import com.example.kudaki.MainActivity;
import com.example.kudaki.R;
import com.example.kudaki.event.EventActivity;
import com.example.kudaki.explore.ExploreActivity;
import com.example.kudaki.model.response.AddressData;
import com.example.kudaki.model.response.ProfileData;
import com.example.kudaki.profile.owner.OwnerFragment;
import com.example.kudaki.profile.tenant.TenantFragment;
import com.example.kudaki.renting.RentalActivity;
import com.example.kudaki.setting.SettingActivity;
import com.google.android.material.bottomnavigation.BottomNavigationView;
import com.google.android.material.tabs.TabLayout;
import com.orhanobut.hawk.Hawk;

import butterknife.BindView;
import butterknife.ButterKnife;

public class ProfileActivity extends AppCompatActivity implements ProfileContract.View {
    @BindView(R.id.profileToolbar)
    Toolbar toolbar;
    @BindView(R.id.profileNav)
    BottomNavigationView bottomNav;
    @BindView(R.id.profileName)
    TextView name;
    @BindView(R.id.profilePhone)
    TextView phone;
    @BindView(R.id.profileTab)
    TabLayout tabLayout;
    @BindView(R.id.profileViewPager)
    ViewPager viewPager;

    String token;
    TabAdapter adapter;

    ProfileContract.Presenter contractPresenter;
    ProfilePresenter profilePresenter;

    ProgressDialog progressDialog;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_profile);
        overridePendingTransition(android.R.anim.fade_in, android.R.anim.fade_out);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);

        Hawk.init(this).build();

        token = Hawk.get("token");

        profilePresenter = new ProfilePresenter(this, token);

        progressDialog = new ProgressDialog(this);

        adapter = new TabAdapter(getSupportFragmentManager());
        adapter.addFragment(new TenantFragment(), "Pengguna");
        adapter.addFragment(new OwnerFragment(), "Pemilik");
        viewPager.setAdapter(adapter);
        tabLayout.setupWithViewPager(viewPager);
    }

    @Override
    protected void onStart() {
        super.onStart();

        contractPresenter.loadProfile();
    }

    @Override
    protected void onResume() {
        super.onResume();

        bottomNav.getMenu().getItem(4).setChecked(true);
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
                    startActivity(new Intent(this, RentalActivity.class));
                    finish();
                    return true;
                case R.id.navProfile:
                    return true;
            }
            return false;
        });
    }

    @Override
    public boolean onCreateOptionsMenu(@NonNull Menu menu) {
        getMenuInflater().inflate(R.menu.profile_menu, menu);
        return super.onCreateOptionsMenu(menu);
    }

    @Override
    public boolean onOptionsItemSelected(@NonNull MenuItem item) {
        switch (item.getItemId()) {
            case R.id.setting:
                Intent setting = new Intent(this, SettingActivity.class);
                setting.putExtra("token", token);
                startActivity(setting);
                return true;
        }

        return super.onOptionsItemSelected(item);
    }

    @Override
    public void setPresenter(ProfileContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }

    @Override
    public void showProfileData(ProfileData data) {
        name.setText(data.getProfile().getFullName());
        phone.setText(data.getProfile().getPhoneNumber());
    }

    @Override
    public void checkAddress(AddressData data) {

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
}
