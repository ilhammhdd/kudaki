package com.example.kudaki.register;

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
import com.example.kudaki.model.user.User;

import butterknife.BindView;
import butterknife.ButterKnife;

public class RegisterActivity extends AppCompatActivity implements RegisterContract.View {
    @BindView(R.id.registerName)
    EditText name;
    @BindView(R.id.registerEmail)
    EditText email;
    @BindView(R.id.registerPhone)
    EditText phone;
    @BindView(R.id.registerPassword)
    EditText password;
    @BindView(R.id.registerConfirm)
    EditText confirm;
    @BindView(R.id.submitRegister)
    Button register;
    @BindView(R.id.registerToolbar)
    Toolbar toolbar;

    RegisterPresenter registerPresenter;
    RegisterContract.Presenter contractPresenter;

    ProgressDialog progressDialog;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_register);
        ButterKnife.bind(this);

        progressDialog = new ProgressDialog(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        // assign presenter to this activity
        registerPresenter = new RegisterPresenter(this);
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
    protected void onResume() {
        super.onResume();

        register.setOnClickListener(view -> {
            if (!password.getText().toString().equals(confirm.getText().toString())) {
                Toast.makeText(this, "Daftar gagal! Password Anda tidak sesuai!", Toast.LENGTH_SHORT).show();
            } else {
                contractPresenter.doRegister(new User(
                        name.getText().toString(),
                        email.getText().toString(),
                        password.getText().toString(),
                        phone.getText().toString()));
            }
        });
    }

    @Override
    public void setPresenter(RegisterContract.Presenter contractPresenter) {
        this.contractPresenter = contractPresenter;
    }

    @Override
    public void showOnRegisterSuccess(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
        NavUtils.navigateUpFromSameTask(this);
    }

    @Override
    public void showOnRegisterFailed(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
    }

    @Override
    public void showLoginActivity() {
        NavUtils.navigateUpFromSameTask(this);
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
    protected void onStop() {
        super.onStop();
    }

    @Override
    protected void onDestroy() {
        super.onDestroy();
    }

    @Override
    public void onBackPressed() {
        NavUtils.navigateUpFromSameTask(this);
    }
}
