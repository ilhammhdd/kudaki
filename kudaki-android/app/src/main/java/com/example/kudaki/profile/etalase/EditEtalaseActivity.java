package com.example.kudaki.profile.etalase;

import android.app.ProgressDialog;
import android.os.Bundle;
import android.view.MenuItem;
import android.view.View;
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

public class EditEtalaseActivity extends AppCompatActivity implements EditEtalaseContract.View {
    @BindView(R.id.editEtalaseToolbar)
    Toolbar toolbar;
    @BindView(R.id.editEtalaseName)
    EditText name;
    @BindView(R.id.editEtalaseDesc)
    EditText desc;
    @BindView(R.id.editEtalasePrice)
    EditText price;
    @BindView(R.id.editEtalaseSave)
    Button btnSave;

    String token;
    ProgressDialog progressDialog;

    EditEtalaseContract.Presenter contractPresenter;
    EditEtalasePresenter presenter;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_edit_etalase);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        Hawk.init(this).build();

        token = Hawk.get("token");

        presenter = new EditEtalasePresenter(this, token);

        progressDialog = new ProgressDialog(this);
    }

    @Override
    protected void onStart() {
        super.onStart();

        name.setText(getIntent().getExtras().getString("name"));
        desc.setText(getIntent().getExtras().getString("description"));
        price.setText(String.valueOf(getIntent().getExtras().getInt("price")));
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

        btnSave.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                contractPresenter.update(
                        getIntent().getExtras().getString("uuid"),
                        name.getText().toString(),
                        desc.getText().toString(),
                        price.getText().toString(),
                        "DAY"
                );
            }
        });
    }

    @Override
    public void setPresenter(EditEtalaseContract.Presenter presenter) {
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
    public void showEditSuccess(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
        NavUtils.navigateUpFromSameTask(this);
    }

    @Override
    public void showEditFailed(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
    }
}
