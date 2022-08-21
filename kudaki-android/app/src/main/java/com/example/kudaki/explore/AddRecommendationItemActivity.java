package com.example.kudaki.explore;

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
import com.orhanobut.hawk.Hawk;

import butterknife.BindView;
import butterknife.ButterKnife;

public class AddRecommendationItemActivity extends AppCompatActivity implements AddRecommendationItemContract.View {
    @BindView(R.id.addRecommendationItemToolbar)
    Toolbar toolbar;
    @BindView(R.id.recommendationItemName)
    EditText name;
    @BindView(R.id.recommendationItemTotal)
    EditText total;
    @BindView(R.id.recommendationItemAdd)
    Button add;

    String token, uuid;

    AddRecommendationItemContract.Presenter contractPresenter;
    AddRecommendationItemPresenter presenter;

    ProgressDialog progressDialog;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_add_recommendation_item);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        Hawk.init(this).build();

        token = Hawk.get("token");

        presenter = new AddRecommendationItemPresenter(this, token, uuid);

        uuid = getIntent().getExtras().getString("uuid");

        progressDialog = new ProgressDialog(this);
    }

    @Override
    protected void onResume() {
        super.onResume();

        add.setOnClickListener(v -> contractPresenter.add(
                name.getText().toString(),
                total.getText().toString()
        ));
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
    public void showAddSuccess(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
        NavUtils.navigateUpFromSameTask(this);
    }

    @Override
    public void showAddFailed(String message) {
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
    public void setPresenter(AddRecommendationItemContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }
}
