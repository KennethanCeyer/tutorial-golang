<h1 align="center">Do It Golang!</h1>
<p align="center">Golang 학습을 위한 예제 코드 모음</p>

### 소개

이 저장소는 Do It Golang! 책에 수록된 모든 예제 코드와 연습 문제들을 관리합니다.
각 챕터별로 디렉토리로 구성되어 있으며, Golang의 기본 문법부터 심화 주제까지 다양한 예제들을 포함하고 있습니다.

### 파일 구조

| 챕터/폴더   | 설명                                                         |
|------------|--------------------------------------------------------------|
| `ch01`     | Go 언어 소개 및 설치, 첫 번째 Go 프로그램 예제 코드 포함      |
| `ch02`     | Go 개발 환경 설정, 환경 변수 등록, 외부 패키지 활용 예제 코드  |
| `ch03`     | 변수, 상수, 자료형, 지역/전역 변수 관련 예제 코드             |
| `ch04`     | 기본 자료형, 자료형 변환, 포인터 사용 예제 코드               |
| `ch05`     | 산술, 비교, 논리, 비트 연산자 및 기타 연산자 예제 코드        |
| `ch06`     | 함수 선언, 다중 반환값, 익명 함수, 클로저 활용 예제 코드      |
| `ch07`     | 조건문(`if`, `switch`), 논리 연산자를 활용한 조건문 예제 코드 |
| `ch08`     | 반복문(`for`), 반복문 제어, 구구단 프로그램 구현 예제 코드    |
| `ch09`     | 배열, 슬라이스, 맵, 구조체 등 자료구조 활용 예제 코드         |
| `ch10`     | 객체지향 개념, 메서드, 인터페이스, 리시버, 상속 관련 예제 코드 |
| `ch11`     | 표준 입출력 및 파일 입출력 관련 예제 코드                     |
| `ch12`     | 오류 처리, `panic` 발생 및 복구, `defer` 예약어 활용 예제 코드 |
| `ch13`     | 동시성 프로그래밍, 고루틴, 뉴스레터 가져오기 예제 코드        |
| `ch14`     | 동시성 제어 기법, 채널, 선택문, 잠금 및 대기 그룹 예제 코드   |
| `ch15`     | 제네릭 개념, 함수와 타입, 인터페이스 활용 예제 코드           |
| `ch16`     | 성능 최적화, 프로파일링 및 프로파일 기반 최적화 예제 코드     |
| `ch17`     | 네트워킹 개념, TCP, UDP, HTTP 네트워킹 관련 예제 코드         |
| `ch18`     | 할 일 관리 애플리케이션 및 도서 정보 관리 웹 애플리케이션 구현 |
| `ch19`     | 단위 테스트, 단언문, 테스트 커버리지, 모킹과 스터빙 예제 코드  |
| `ch20`     | 리팩터링 개념, 코드 개선 기법, 코드 구조 변경 예제 코드       |
| `appendix` | 부록: `godoc`, `context`, `cgo` 패키지 활용 관련 예제 코드 |
| `extra`    | 고차 함수 등 그 밖에 예제 코드                            |


### 실행 환경

실행 환경은 아래 기준으로 테스트가 되었습니다.

| 항목            | 내용                                         |
|-----------------|----------------------------------------------|
| Golang 버전     | Golang 1.24.0          |
| 운영체제        | Linux (Debian), macOS, Windows                      |
| 기타 도구       | Git, VSCode (권장)                            |


### 사용법

1. 저장소 클론

Git 클라이언트를 통해 저장소를 로컬에 클론합니다.

```bash
git clone https://github.com/KennethanCeyer/tutorial-golang.git
```

2. 디렉토리 이동 및 코드 확인

원하는 챕터의 디렉토리로 이동하여 예제 코드를 확인할 수 있습니다.

```bash
# chapter 10의 file_io로 접근하는 경우
cd ch10/file_io
```

3. 코드 실행

Golang이 설치되어 있다면, 아래와 같이 예제 코드를 실행할 수 있습니다.

```bash
go run file_io.go
```

4. 빌드 및 테스트

코드를 빌드하거나 테스트할 수 있습니다.

```bash
go build go test
```

### 라이선스

이 저장소에 포함된 코드는 이지스퍼블리싱의 Do It Golang! 출판용 라이선스를 따릅니다.

- 학업적 목적, 블로그 작성, 교육 과제 등 학습 및 교육 목적으로 자유롭게 사용할 수 있습니다.
- 무단 수정하여 다른 출판물에 활용 등 상업적 목적의 사용은 불가합니다.
   - 단, 강의 및 교육 목적으로 상업적 환경에서 사용되는 것은 허용됩니다.
