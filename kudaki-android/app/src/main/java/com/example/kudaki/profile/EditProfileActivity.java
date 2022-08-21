package com.example.kudaki.profile;

import android.app.ProgressDialog;
import android.os.Bundle;
import android.view.MenuItem;
import android.widget.Button;
import android.widget.EditText;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.core.app.NavUtils;

import com.example.kudaki.R;
import com.example.kudaki.model.response.ProfileData;
import com.orhanobut.hawk.Hawk;

import butterknife.BindView;
import butterknife.ButterKnife;

public class EditProfileActivity extends AppCompatActivity implements EditProfileContract.View {
    @BindView(R.id.editProfileToolbar)
    Toolbar toolbar;
    @BindView(R.id.editName)
    EditText name;
    @BindView(R.id.editPhone)
    EditText phone;
    @BindView(R.id.btnSaveProfile)
    Button btnSave;

    String token;

    EditProfileContract.Presenter contractPresenter;
    EditProfilePresenter presenter;

    ProgressDialog progressDialog;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_edit_profile);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        Hawk.init(this).build();

        token = Hawk.get("token");

        presenter = new EditProfilePresenter(this, token);

        progressDialog = new ProgressDialog(this);
    }

    @Override
    protected void onStart() {
        super.onStart();

        contractPresenter.loadProfile();
    }

    @Override
    protected void onResume() {
        super.onResume();

        btnSave.setOnClickListener(v -> contractPresenter.update(
                name.getText().toString(),
                phone.getText().toString()));
    }

    @Override
    public void setPresenter(EditProfileContract.Presenter presenter) {
        this.contractPresenter = presenter;
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
    public void showProfileData(ProfileData data) {
        name.setText(data.getProfile().getFullName());
        phone.setText(data.getProfile().getPhoneNumber());
    }

    @Override
    public void showEditSuccess(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
        NavUtils.navigateUpFromSameTask(this);
    }

    @Override
    public void showEditFailed(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
    }
}
