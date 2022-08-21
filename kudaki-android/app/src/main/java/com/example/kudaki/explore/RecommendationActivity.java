package com.example.kudaki.explore;

import android.app.AlertDialog;
import android.app.ProgressDialog;
import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.MenuItem;
import android.view.View;
import android.widget.EditText;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.core.app.NavUtils;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.kudaki.R;
import com.example.kudaki.adapter.RecommendationAdapter;
import com.example.kudaki.model.request.RecommendationItem;
import com.example.kudaki.model.request.RecommendationRequest;
import com.example.kudaki.model.response.RecommendationData;
import com.example.kudaki.model.response.RecommendedGear;
import com.google.android.material.floatingactionbutton.FloatingActionButton;
import com.orhanobut.hawk.Hawk;

import java.util.ArrayList;

import butterknife.BindView;
import butterknife.ButterKnife;

public class RecommendationActivity extends AppCompatActivity implements RecommendationContract.View {
    @BindView(R.id.recommendationToolbar)
    Toolbar toolbar;
    @BindView(R.id.rvRecommendation)
    RecyclerView recyclerView;
    @BindView(R.id.recommendationAdd)
    FloatingActionButton fabAdd;

    String token, uuid;
    RecommendationAdapter adapter;

    RecommendationContract.Presenter contractPresenter;
    RecommendationPresenter presenter;

    ArrayList<RecommendedGear> list;

    ProgressDialog progressDialog;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_recommendation);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        Hawk.init(this).build();

        token = Hawk.get("token");
        uuid = getIntent().getExtras().getString("uuid");

        presenter = new RecommendationPresenter(this, token, uuid);

        progressDialog = new ProgressDialog(this);
    }

    @Override
    protected void onStart() {
        super.onStart();

        contractPresenter.loadRecommendation();
    }

    @Override
    protected void onResume() {
        super.onResume();

        fabAdd.setOnClickListener(v -> {
            AlertDialog.Builder builder = new AlertDialog.Builder(v.getContext(), R.style.CustomDialogTheme);

            LayoutInflater inflater = LayoutInflater.from(v.getContext());

            View view = inflater.inflate(R.layout.dialog_recommend, null);
            EditText name = view.findViewById(R.id.recommendName);
            EditText amount = view.findViewById(R.id.recommendAmount);

            builder.setView(view);
            builder.setTitle("Tambah Rekomendasi");
            builder.setPositiveButton("Tambah", (dialog, which) -> {
                RecommendationRequest request = new RecommendationRequest();
                RecommendationItem item = new RecommendationItem();
                item.setType(name.getText().toString());
                item.setTotal(Integer.parseInt(amount.getText().toString()));

                ArrayList<RecommendationItem> items = new ArrayList<>();
                items.add(item);

                request.setUuid(uuid);
                request.setItems(items);
                contractPresenter.add(request);
            });
            builder.setNegativeButton("Batal", (dialog, which) -> dialog.dismiss());
            builder.show();
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
    public void showAddSuccess(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
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
    public void showData(RecommendationData data) {
        if (data.getRecommendedGears() == null) {
            Toast.makeText(this, "Belum ada rekomendasi alat", Toast.LENGTH_SHORT).show();
        } else {
            adapter = new RecommendationAdapter(this, data.getRecommendedGears());
            adapter.notifyDataSetChanged();
            adapter.setToken(token);
            recyclerView.setLayoutManager(new LinearLayoutManager(this));
            recyclerView.setAdapter(adapter);
        }
    }

    @Override
    public void setPresenter(RecommendationContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }
}
