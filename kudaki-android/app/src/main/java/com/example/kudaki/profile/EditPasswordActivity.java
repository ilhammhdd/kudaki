package com.example.kudaki.profile;

import android.app.ProgressDialog;
import android.os.Bundle;
import android.widget.Button;
import android.widget.EditText;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.core.app.NavUtils;

import com.example.kudaki.R;
import com.orhanobut.hawk.Hawk;

import butterknife.BindView;
import butterknife.ButterKnife;

public class EditPasswordActivity extends AppCompatActivity implements EditPasswordContract.View{
    @BindView(R.id.editPasswordToolbar)
    Toolbar toolbar;
    @BindView(R.id.oldPassword)
    EditText oldPwd;
    @BindView(R.id.newPassword)
    EditText newPwd;
    @BindView(R.id.confirmPassword)
    EditText confirmPwd;
    @BindView(R.id.btnSavePassword)
    Button btnSave;

    String token;

    EditPasswordContract.Presenter contractPresenter;
    EditPasswordPresenter presenter;

    ProgressDialog progressDialog;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_edit_password);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        Hawk.init(this).build();

        token = Hawk.get("token");

        presenter = new EditPasswordPresenter(this, token);

        progressDialog = new ProgressDialog(this);
    }

    @Override
    protected void onResume() {
        super.onResume();

        btnSave.setOnClickListener(v -> {
            if (newPwd.getText().toString().equals(confirmPwd.getText().toString())) {
                contractPresenter.update(oldPwd.getText().toString(), newPwd.getText().toString());
            } else {
                Toast.makeText(v.getContext(), "Password tidak sesuai", Toast.LENGTH_SHORT).show();
            }
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
    public void showEditSuccess(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
        NavUtils.navigateUpFromSameTask(this);
    }

    @Override
    public void showEditFailed(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
    }

    @Override
    public void setPresenter(EditPasswordContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }
}
