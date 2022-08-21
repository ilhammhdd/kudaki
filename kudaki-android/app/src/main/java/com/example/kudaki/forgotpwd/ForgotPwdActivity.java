package com.example.kudaki.forgotpwd;

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

import butterknife.BindView;
import butterknife.ButterKnife;

public class ForgotPwdActivity extends AppCompatActivity implements ForgotPwdContract.View {
    @BindView(R.id.btnSendEmail)
    Button btnSendEmail;
    @BindView(R.id.inForgotEmail)
    EditText forgotEmail;
    @BindView(R.id.forgotToolbar)
    Toolbar toolbar;

    ForgotPwdPresenter forgotPwdPresenter;
    ForgotPwdContract.Presenter contractPresenter;

    ProgressDialog progressDialog;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_forgot_pwd);
        ButterKnife.bind(this);

        progressDialog = new ProgressDialog(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        forgotPwdPresenter = new ForgotPwdPresenter(this);
    }

    @Override
    protected void onResume() {
        super.onResume();

        btnSendEmail.setOnClickListener(view -> {
            String email = forgotEmail.getText().toString();
            contractPresenter.doSendEmail(email);
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
    public void showSendSuccess(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
        NavUtils.navigateUpFromSameTask(this);
    }

    @Override
    public void showSendFailed(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
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
    public void setPresenter(ForgotPwdContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }
}
