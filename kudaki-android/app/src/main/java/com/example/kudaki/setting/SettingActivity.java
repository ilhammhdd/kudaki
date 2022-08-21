package com.example.kudaki.setting;

import android.content.Intent;
import android.os.Bundle;
import android.view.MenuItem;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.cardview.widget.CardView;
import androidx.core.app.NavUtils;

import com.example.kudaki.R;
import com.example.kudaki.login.LoginActivity;
import com.example.kudaki.profile.EditPasswordActivity;
import com.example.kudaki.profile.EditProfileActivity;
import com.orhanobut.hawk.Hawk;

import butterknife.BindView;
import butterknife.ButterKnife;

public class SettingActivity extends AppCompatActivity implements SettingContract.View {
    @BindView(R.id.settingToolbar)
    Toolbar toolbar;
    @BindView(R.id.cardEditPassword)
    CardView editPassword;
    @BindView(R.id.cardEditProfile)
    CardView editProfile;
    @BindView(R.id.cardLogout)
    CardView logout;

    SettingPresenter settingPresenter;
    SettingContract.Presenter contractPresenter;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_setting);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        // assign presenter to this activity
        settingPresenter = new SettingPresenter(this);
    }

    @Override
    protected void onResume() {
        super.onResume();

        editPassword.setOnClickListener(view -> {
            Intent password = new Intent(SettingActivity.this, EditPasswordActivity.class);
            startActivity(password);
        });

        editProfile.setOnClickListener(view -> {
            Intent password = new Intent(SettingActivity.this, EditProfileActivity.class);
            startActivity(password);
        });

        logout.setOnClickListener(view -> {
            contractPresenter.doLogout(this);
        });
    }

    @Override
    public void setPresenter(SettingContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }

    @Override
    public void showLogoutSuccess(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
        Intent login = new Intent(this, LoginActivity.class);
        login.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK | Intent.FLAG_ACTIVITY_CLEAR_TASK);

        Hawk.deleteAll();

        startActivity(login);
        finish();
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
}
