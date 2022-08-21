package com.example.kudaki.reset;

import android.app.ProgressDialog;
import android.content.Intent;
import android.net.Uri;
import android.os.Bundle;
import android.view.MenuItem;
import android.widget.Button;
import android.widget.EditText;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.core.app.NavUtils;

import com.example.kudaki.R;
import com.example.kudaki.login.LoginActivity;

import butterknife.BindView;
import butterknife.ButterKnife;

public class ResetActivity extends AppCompatActivity implements ResetContract.View {
    @BindView(R.id.resetSubmit)
    Button btnSubmit;
    @BindView(R.id.resetPassword)
    EditText newPwd;
    @BindView(R.id.resetConfirm)
    EditText confrimPwd;
    @BindView(R.id.resetToolbar)
    Toolbar toolbar;

    Intent intent;
    Uri data;
    String token;

    ResetContract.Presenter contractPresenter;
    ResetPresenter presenter;

    ProgressDialog progressDialog;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_reset);
        ButterKnife.bind(this);

        progressDialog = new ProgressDialog(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        presenter = new ResetPresenter(this, this);

        intent = getIntent();
        data = intent.getData();
    }

    @Override
    protected void onStart() {
        super.onStart();

         // get reset_token value
         token = data.getQueryParameter("reset_token");
    }

    @Override
    protected void onResume() {
        super.onResume();

        btnSubmit.setOnClickListener(v -> {
            String pass = newPwd.getText().toString();
            String confirm = confrimPwd.getText().toString();
            if (pass.equals(confirm)) {
                contractPresenter.doReset(token, newPwd.getText().toString());
            } else {
                Toast.makeText(ResetActivity.this, "Password Tidak Sesuai!", Toast.LENGTH_SHORT).show();
            }
        });
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
    public void setPresenter(ResetContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }

    @Override
    public void showResetSuccess(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
        Intent login = new Intent(this, LoginActivity.class);
        startActivity(login);
        finish();
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
    public void showResetFailed(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
    }
}
