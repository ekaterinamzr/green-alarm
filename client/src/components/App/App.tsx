import { FC, lazy, Suspense } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { ConfigProvider, Spin } from "antd";
import ruRU from "antd/locale/ru_RU";

const IncidentsPage = lazy(() => import("pages/incidents"));
const MapPage = lazy(() => import("pages/map"));
const SignInPage = lazy(() => import("pages/signin"));
const SignUpPage = lazy(() => import("pages/signup"));

const App: FC = () => {
  return (
    <ConfigProvider locale={ruRU}>
      <BrowserRouter>
        <Suspense fallback={<Spin />}>
          <Routes>
            <Route path="/incidents/" element={<IncidentsPage />} />
            <Route path="/map/" element={<MapPage />} />
            <Route path="/signin/" element={<SignInPage />} />
            <Route path="/signup/" element={<SignUpPage />} />
          </Routes>
        </Suspense>
      </BrowserRouter>
    </ConfigProvider>
  );
};

export { App };
