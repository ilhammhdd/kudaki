package com.example.kudaki;

import android.content.Intent;
import android.os.Bundle;
import android.os.Handler;

import androidx.appcompat.app.AppCompatActivity;

import com.example.kudaki.login.LoginActivity;
import com.orhanobut.hawk.Hawk;

public class SplashActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.splash_screen);

        Hawk.init(this).build();

        Boolean isIntroOpened = Hawk.get("isIntroOpened");
        String token = Hawk.get("token");

        final Handler handler = new Handler();
        handler.postDelayed(() -> {
            if (isIntroOpened == null) {
                Intent tutorial = new Intent(getApplicationContext(), TutorialActivity.class);
                startActivity(tutorial);
                finish();
            } else if (token == null) {
                Intent login = new Intent(getApplicationContext(), LoginActivity.class);
                startActivityForResult(login, 1);
                finish();
            } else {
                Intent home = new Intent(getApplicationContext(), MainActivity.class);
                startActivity(home);
                finish();
            }
        }, 3000);
    }
}
