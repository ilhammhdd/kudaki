package com.example.kudaki.explore;

import android.content.Intent;
import android.net.Uri;
import android.os.Bundle;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AlertDialog;
import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.core.app.NavUtils;

import com.bumptech.glide.Glide;
import com.example.kudaki.R;
import com.google.android.material.floatingactionbutton.FloatingActionButton;

import butterknife.BindView;
import butterknife.ButterKnife;

public class MountainActivity extends AppCompatActivity {
    @BindView(R.id.mountainToolbar)
    Toolbar toolbar;
    @BindView(R.id.mountainName)
    TextView name;
    @BindView(R.id.mountainDesc)
    TextView description;
    @BindView(R.id.mountainHeight)
    TextView height;
    @BindView(R.id.mountainDifficulty)
    TextView difficulty;
    @BindView(R.id.mountainImage)
    ImageView image;
    @BindView(R.id.fabGoogleMap)
    FloatingActionButton floatingActionButton;

    String uuid;
    Bundle bundle;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_mountain);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        bundle = getIntent().getExtras();
        toolbar.setTitle(bundle.getString("name"));

        uuid = bundle.getString("uuid");
    }

    @Override
    protected void onStart() {
        super.onStart();

        Glide.with(this)
                .load(bundle.getString("photo"))
                .into(image);
        name.setText(bundle.getString("name"));
        description.setText(bundle.getString("description"));
        difficulty.setText(String.valueOf(bundle.getDouble("difficulty", 0)));
        height.setText(bundle.getInt("height") + " Mdpl");
    }

    @Override
    protected void onResume() {
        super.onResume();

        floatingActionButton.setOnClickListener(v -> {
            Uri gmmIntentUri = Uri.parse(
                    "geo:"+ bundle.getDouble("latitude") +"," + bundle.getDouble("latitude"));
            Intent mapIntent = new Intent(Intent.ACTION_VIEW, gmmIntentUri);
            mapIntent.setPackage("com.google.android.apps.maps");
            if (mapIntent.resolveActivity(getPackageManager()) != null) {
                startActivity(mapIntent);
            }
        });

        description.setOnClickListener(v -> {
            AlertDialog.Builder builder = new AlertDialog.Builder(v.getContext(), R.style.CustomDialogTheme);

            builder.setTitle("Deskripsi Gunung");
            builder.setMessage(bundle.getString("description"));
            builder.setNeutralButton("Tutup", (dialog, which) -> dialog.dismiss());
            builder.show();
        });
    }

    @Override
    public boolean onCreateOptionsMenu(@NonNull Menu menu) {
        getMenuInflater().inflate(R.menu.explore_menu, menu);
        return super.onCreateOptionsMenu(menu);
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        switch (item.getItemId()) {
            case android.R.id.home:
                NavUtils.navigateUpFromSameTask(this);
                return true;
            case R.id.recommendation:
                Intent intent = new Intent(this, RecommendationActivity.class);
                intent.putExtra("uuid", uuid);
                startActivity(intent);
                return true;
        }
        return super.onOptionsItemSelected(item);
    }
}
